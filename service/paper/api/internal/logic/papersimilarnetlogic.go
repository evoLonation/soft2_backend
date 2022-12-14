package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"soft2_backend/service/paper/database"
	"strconv"

	"soft2_backend/service/paper/api/internal/svc"
	"soft2_backend/service/paper/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaperSimilarNetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaperSimilarNetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaperSimilarNetLogic {
	return &PaperSimilarNetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var nodesSimilar []types.PaperNodeJSON
var edgesSimilar []types.EdgeJSON
var maxYearSimilar int
var minYearSimilar int
var maxCitationSimilar int
var minCitationSimilar int
var nodeMapSimilar map[string]int

func (l *PaperSimilarNetLogic) PaperSimilarNet(req *types.PaperSimilarNetRequest) (resp *types.PaperSimilarNetResponse, err error) {
	// todo: add your logic here and delete this line
	nodesSimilar = make([]types.PaperNodeJSON, 0)
	edgesSimilar = make([]types.EdgeJSON, 0)
	maxYearSimilar = 0
	minYearSimilar = 3000
	maxCitationSimilar = 0
	minCitationSimilar = 1000000
	nodeMapSimilar = make(map[string]int, 0)

	var thisPaperBuf bytes.Buffer
	thisPaperQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"id": req.Id,
			},
		},
	}
	if err := json.NewEncoder(&thisPaperBuf).Encode(thisPaperQuery); err != nil {
		log.Printf("encode query error: %v", err)
	}
	thisPaperRes := database.SearchPaper(thisPaperBuf)

	thisPaperSource := thisPaperRes["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})

	authors := NilHandler(thisPaperSource["authors"], "list").([]interface{})
	var author string
	if len(authors) == 0 {
		author = ""
	} else {
		author = NilHandler(authors[0].(map[string]interface{})["name"], "string").(string)
	}
	majorNode := types.PaperNodeJSON{
		Id:    req.Id,
		Label: author + strconv.Itoa(NilHandler(thisPaperSource["year"], "int").(int)),
		Size:  NilHandler(thisPaperSource["n_citation"], "int").(int),
		Type:  "major",
		Style: types.StyleJSON{
			Fill: strconv.Itoa(NilHandler(thisPaperSource["year"], "int").(int)),
		},
		Info: types.InfoJSON{
			Id:       req.Id,
			Title:    NilHandler(thisPaperSource["title"], "string").(string),
			Abstract: NilHandler(thisPaperSource["abstract"], "string").(string),
			Authors:  GetPaperAuthors(thisPaperSource),
			Year:     NilHandler(thisPaperSource["year"], "int").(int),
		},
	}
	nodesSimilar = append(nodesSimilar, majorNode)
	nodeMapSimilar[req.Id] = 0
	UpdateMaxMin(&maxYearSimilar, &minYearSimilar, majorNode.Info.Year)
	UpdateMaxMin(&maxCitationSimilar, &minCitationSimilar, majorNode.Size)

	references := NilHandler(thisPaperSource["relateds"], "list").([]interface{})
	referenceIds := make([]string, 0)
	for _, reference := range references {
		_, ok := nodeMapSimilar[reference.(string)]
		if ok {
			continue
		}
		referenceIds = append(referenceIds, reference.(string))
	}
	log.Println(referenceIds)
	log.Println(nodeMapSimilar)

	DFSSimilar(referenceIds, majorNode, 0)

	log.Printf("maxCitation: %d, minCitation: %d", maxCitationSimilar, minCitationSimilar)
	for i, node := range nodesSimilar {
		log.Printf("node %d: %d", i, node.Size)
		nodesSimilar[i].Size = GetSize(node.Size, maxCitationSimilar, minCitationSimilar)
		nodesSimilar[i].Style.Fill = GetColor(GetD(node.Info.Year, maxYearSimilar, minYearSimilar))
	}

	resp = &types.PaperSimilarNetResponse{
		Nodes: nodesSimilar,
		Edges: edgesSimilar,
	}
	return resp, nil
}

func DFSSimilar(referenceIds []string, fatherNode types.PaperNodeJSON, level int) {
	if level == 3 {
		return
	}
	log.Println(level)

	var referenceBuf bytes.Buffer
	referenceQuery := map[string]interface{}{
		"ids": referenceIds,
	}
	if err := json.NewEncoder(&referenceBuf).Encode(referenceQuery); err != nil {
		log.Printf("encode query error: %v", err)
	}
	referenceRes := database.MgetPaper(referenceBuf)

	papers := NilHandler(referenceRes["docs"], "list").([]interface{})
	var levelPaperCnt = 0
	for _, paper := range papers {
		if paper.(map[string]interface{})["found"].(bool) == false {
			continue
		}
		if (levelPaperCnt >= 4 && level == 0) ||
			(levelPaperCnt >= 3 && level == 1) ||
			(levelPaperCnt >= 3 && level == 2) {
			break
		}
		levelPaperCnt++
		source := paper.(map[string]interface{})["_source"].(map[string]interface{})
		authors := NilHandler(source["authors"], "list").([]interface{})
		var author string
		if len(authors) == 0 {
			author = ""
		} else {
			author = NilHandler(authors[0].(map[string]interface{})["name"], "string").(string)
		}
		node := types.PaperNodeJSON{
			Id:    NilHandler(source["id"], "string").(string),
			Label: author + strconv.Itoa(NilHandler(source["year"], "int").(int)),
			Size:  NilHandler(source["n_citation"], "int").(int),
			Style: types.StyleJSON{
				Fill: strconv.Itoa(NilHandler(source["year"], "int").(int)),
			},
			Info: types.InfoJSON{
				Id:       NilHandler(source["id"], "string").(string),
				Title:    NilHandler(source["title"], "string").(string),
				Abstract: NilHandler(source["abstract"], "string").(string),
				Authors:  GetPaperAuthors(source),
				Year:     NilHandler(source["year"], "int").(int),
			},
		}
		nodesSimilar = append(nodesSimilar, node)
		nodeMapSimilar[node.Id] = len(nodesSimilar) - 1
		UpdateMaxMin(&maxYearSimilar, &minYearSimilar, node.Info.Year)
		UpdateMaxMin(&maxCitationSimilar, &minCitationSimilar, node.Size)

		edgesSimilar = append(edgesSimilar, types.EdgeJSON{
			Source: fatherNode.Id,
			Target: node.Id,
		})

		references := NilHandler(source["relateds"], "list").([]interface{})
		referenceIds = make([]string, 0)
		log.Println(nodeMapSimilar)
		for _, reference := range references {
			_, ok := nodeMapSimilar[reference.(string)]
			if ok {
				continue
			}
			referenceIds = append(referenceIds, reference.(string))
		}
		log.Println(referenceIds)

		DFSSimilar(referenceIds, node, level+1)
	}
}
