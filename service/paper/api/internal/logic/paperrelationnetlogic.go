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

	majorNode := types.PaperNodeJSON{
		Id:    req.Id,
		Label: thisPaperSource["authors"].([]interface{})[0].(map[string]interface{})["name"].(string) + strconv.Itoa(NilHandler(thisPaperSource["year"], "int").(int)),
		Size:  NilHandler(thisPaperSource["n_citation"], "int").(int),
		Type:  "major",
		Style: types.StyleJSON{
			Fill: strconv.Itoa(NilHandler(thisPaperSource["year"], "int").(int)),
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

	references := NilHandler(thisPaperSource["references"], "list").([]interface{})
	referenceIds := make([]string, 0)
	for _, reference := range references {
		referenceIds = append(referenceIds, reference.(string))
	}

	DFS(referenceIds, majorNode, 0)

	for _, node := range nodes {
		node.Size = GetSize(node.Size, maxCitation, minCitation)
		node.Style.Fill = GetColor(GetD(node.Info.Year, maxYear, minYear))
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

	papers := NilHandler(referenceRes["docs"], "list").([]interface{})
	for _, paper := range papers {
		source := paper.(map[string]interface{})["_source"].(map[string]interface{})
		node := types.PaperNodeJSON{
			Id:    source["id"].(string),
			Label: source["authors"].([]interface{})[0].(map[string]interface{})["name"].(string) + strconv.Itoa(NilHandler(source["year"], "int").(int)),
			Size:  NilHandler(source["n_citation"], "int").(int),
			Style: types.StyleJSON{
				Fill: strconv.Itoa(NilHandler(source["year"], "int").(int)),
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

		edges = append(edges, types.EdgeJSON{
			Source: fatherNode.Id,
			Target: node.Id,
		})

		references := NilHandler(source["references"], "list").([]interface{})
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
