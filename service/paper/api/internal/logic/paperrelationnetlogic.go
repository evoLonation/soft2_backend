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

type PaperRelationNetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaperRelationNetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaperRelationNetLogic {
	return &PaperRelationNetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var nodesRelation []types.PaperNodeJSON
var edgesRelation []types.EdgeJSON
var maxYearRelation int
var minYearRelation int
var maxCitationRelation int
var minCitationRelation int
var nodeMapRelation map[string]int

func (l *PaperRelationNetLogic) PaperRelationNet(req *types.PaperRelationNetRequest) (resp *types.PaperRelationNetResponse, err error) {
	// todo: add your logic here and delete this line
	nodesRelation = make([]types.PaperNodeJSON, 0)
	edgesRelation = make([]types.EdgeJSON, 0)
	maxYearRelation = 0
	minYearRelation = 3000
	maxCitationRelation = 0
	minCitationRelation = 1000000
	nodeMapRelation = make(map[string]int, 0)
	log.Println(nodeMapRelation)

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
	nodesRelation = append(nodesRelation, majorNode)
	nodeMapRelation[req.Id] = 0
	UpdateMaxMin(&maxYearRelation, &minYearRelation, majorNode.Info.Year)
	UpdateMaxMin(&maxCitationRelation, &minCitationRelation, majorNode.Size)

	references := NilHandler(thisPaperSource["references"], "list").([]interface{})
	referenceIds := make([]string, 0)
	for _, reference := range references {
		if len(referenceIds) >= 5 {
			break
		}
		_, ok := nodeMapRelation[reference.(string)]
		if ok {
			continue
		}
		referenceIds = append(referenceIds, reference.(string))
	}

	DFSRelation(referenceIds, majorNode, 0)

	log.Printf("maxCitation: %d, minCitation: %d", maxCitationRelation, minCitationRelation)
	for i, node := range nodesRelation {
		log.Printf("node %d: %d", i, node.Size)
		nodesRelation[i].Size = GetSize(node.Size, maxCitationRelation, minCitationRelation)
		nodesRelation[i].Style.Fill = GetColor(GetD(node.Info.Year, maxYearRelation, minYearRelation))
	}

	resp = &types.PaperRelationNetResponse{
		Nodes: nodesRelation,
		Edges: edgesRelation,
	}
	return resp, nil
}

func DFSRelation(referenceIds []string, fatherNode types.PaperNodeJSON, level int) {
	if level == 3 {
		return
	}

	var referenceBuf bytes.Buffer
	referenceQuery := map[string]interface{}{
		"ids": referenceIds,
	}
	if err := json.NewEncoder(&referenceBuf).Encode(referenceQuery); err != nil {
		log.Printf("encode query error: %v", err)
	}
	referenceRes := database.MgetPaper(referenceBuf)

	papers := NilHandler(referenceRes["docs"], "list").([]interface{})
	for _, paper := range papers {
		if paper.(map[string]interface{})["found"].(bool) == false {
			continue
		}
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
		nodesRelation = append(nodesRelation, node)
		nodeMapRelation[node.Id] = len(nodesRelation) - 1
		UpdateMaxMin(&maxYearRelation, &minYearRelation, node.Info.Year)
		UpdateMaxMin(&maxCitationRelation, &minCitationRelation, node.Size)

		edgesRelation = append(edgesRelation, types.EdgeJSON{
			Source: fatherNode.Id,
			Target: node.Id,
		})

		references := NilHandler(source["references"], "list").([]interface{})
		referenceIds = make([]string, 0)
		for _, reference := range references {
			if (len(referenceIds) >= 5 && level == 0) ||
				(len(referenceIds) >= 4 && level == 1) ||
				(len(referenceIds) >= 3 && level == 2) {
				break
			}
			_, ok := nodeMapRelation[reference.(string)]
			if ok {
				continue
			}
			referenceIds = append(referenceIds, reference.(string))
		}

		DFSRelation(referenceIds, node, level+1)
	}
}
