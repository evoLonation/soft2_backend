package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"soft2_backend/service/paper/database"

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
var edges []types.PaperEdgeJSON
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

	majorNode := types.PaperNodeJSON{
		Id:    req.Id,
		Label: thisPaperSource["authors"].([]interface{})[0].(map[string]interface{})["name"].(string) + NilHandler(thisPaperSource["year"], "int").(string),
		Size:  NilHandler(thisPaperSource["n_citation"], "int").(int),
		Type:  "major",
		Style: types.StyleJSON{
			Fill: NilHandler(thisPaperSource["year"], "int").(string),
		},
		Info: types.InfoJSON{
			Id:       req.Id,
			Title:    thisPaperSource["title"].(string),
			Abstract: NilHandler(thisPaperSource["abstract"], "string").(string),
			Authors:  GetPaperAuthors(thisPaperSource),
			Year:     NilHandler(thisPaperSource["year"], "int").(int),
		},
	}
	nodes = append(nodes, majorNode)
	UpdateMaxMin(&maxYear, &minYear, majorNode.Info.Year)
	UpdateMaxMin(&maxCitation, &minCitation, majorNode.Size)

	references := thisPaperSource["references"].([]interface{})
	referenceIds := make([]string, 0)
	for _, reference := range references {
		referenceIds = append(referenceIds, reference.(string))
	}

	DFS(referenceIds, majorNode, 0)

	for _, node := range nodes {
		node.Size = GetSize(node.Size)
		node.Style.Fill = GetColor(GetD(node.Info.Year))
	}

	resp = &types.PaperRelationNetResponse{
		Nodes: nodes,
		Edges: edges,
	}
	return resp, nil
}

func DFS(referenceIds []string, fatherNode types.PaperNodeJSON, level int) {
	if level == 4 {
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

	papers := referenceRes["docs"].([]interface{})
	for _, paper := range papers {
		source := paper.(map[string]interface{})["_source"].(map[string]interface{})
		node := types.PaperNodeJSON{
			Id:    source["id"].(string),
			Label: source["authors"].([]interface{})[0].(map[string]interface{})["name"].(string) + NilHandler(source["year"], "int").(string),
			Size:  NilHandler(source["n_citation"], "int").(int),
			Style: types.StyleJSON{
				Fill: NilHandler(source["year"], "int").(string),
			},
			Info: types.InfoJSON{
				Id:       source["id"].(string),
				Title:    source["title"].(string),
				Abstract: NilHandler(source["abstract"], "string").(string),
				Authors:  GetPaperAuthors(source),
				Year:     NilHandler(source["year"], "int").(int),
			},
		}
		nodes = append(nodes, node)
		UpdateMaxMin(&maxYear, &minYear, node.Info.Year)
		UpdateMaxMin(&maxCitation, &minCitation, node.Size)

		edges = append(edges, types.PaperEdgeJSON{
			Source: fatherNode.Id,
			Target: node.Id,
		})

		references := source["references"].([]interface{})
		referenceIds = make([]string, 0)
		for _, reference := range references {
			referenceIds = append(referenceIds, reference.(string))
		}

		DFS(referenceIds, node, level+1)
	}
}

func GetPaperAuthors(paper map[string]interface{}) []types.AuthorJSON {
	authors := make([]types.AuthorJSON, 0)
	for _, author := range paper["authors"].([]interface{}) {
		hasId := false
		if author.(map[string]interface{})["id"] != nil {
			hasId = true
		}
		authors = append(authors, types.AuthorJSON{
			Name:  author.(map[string]interface{})["name"].(string),
			Id:    NilHandler(author.(map[string]interface{})["id"], "string").(string),
			HasId: hasId,
		})
	}
	return authors
}

func UpdateMaxMin(max *int, min *int, d int) {
	if d > *max {
		*max = d
	}
	if d < *min {
		*min = d
	}
}

func GetSize(NCitation int) int {
	return int((float64(NCitation-minCitation)/float64(maxCitation-minCitation) + 1) * 20)
}

func GetD(year int) int {
	return int((float64(year-minYear) / float64(maxYear-minYear)) * 10)
}

func GetColor(d int) string {
	switch d {
	case 0:
		return "#1C1C1C"
	case 1:
		return "#363636"
	case 2:
		return "#4F4F4F"
	case 3:
		return "#696969"
	case 4:
		return "#828282"
	case 5:
		return "#9C9C9C"
	case 6:
		return "#B5B5B5"
	case 7:
		return "#CFCFCF"
	case 8:
		return "#E8E8E8"
	default:
		return "#FFFFFF"
	}
}
