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

var nodes []types.PaperNodeJSON
var edges []types.EdgeJSON
var maxYear = 0
var minYear = 3000
var maxCitation = 0
var minCitation = 1000000

func (l *PaperRelationNetLogic) PaperRelationNet(req *types.PaperRelationNetRequest) (resp *types.PaperRelationNetResponse, err error) {
	// todo: add your logic here and delete this line
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
	nodes = append(nodes, majorNode)
	UpdateMaxMin(&maxYear, &minYear, majorNode.Info.Year)
	UpdateMaxMin(&maxCitation, &minCitation, majorNode.Size)

	references := NilHandler(thisPaperSource["references"], "list").([]interface{})
	referenceIds := make([]string, 0)
	for _, reference := range references {
		referenceIds = append(referenceIds, reference.(string))
	}

	DFSRelation(referenceIds, majorNode, 0)

	for i, node := range nodes {
		nodes[i].Size = GetSize(node.Size, maxCitation, minCitation)
		nodes[i].Style.Fill = GetColor(GetD(node.Info.Year, maxYear, minYear))
	}

	resp = &types.PaperRelationNetResponse{
		Nodes: nodes,
		Edges: edges,
	}
	return resp, nil
}

func DFSRelation(referenceIds []string, fatherNode types.PaperNodeJSON, level int) {
	if level == 1 {
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
		log.Println(source)
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
		nodes = append(nodes, node)
		UpdateMaxMin(&maxYear, &minYear, node.Info.Year)
		UpdateMaxMin(&maxCitation, &minCitation, node.Size)

		edges = append(edges, types.EdgeJSON{
			Source: fatherNode.Id,
			Target: node.Id,
		})

		references := NilHandler(source["references"], "list").([]interface{})
		referenceIds = make([]string, 0)
		for _, reference := range references {
			referenceIds = append(referenceIds, reference.(string))
		}

		DFSRelation(referenceIds, node, level+1)
	}
}
