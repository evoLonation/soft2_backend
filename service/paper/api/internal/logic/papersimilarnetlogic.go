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

func (l *PaperSimilarNetLogic) PaperSimilarNet(req *types.PaperSimilarNetRequest) (resp *types.PaperSimilarNetResponse, err error) {
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
		Id: req.Id,
		Label: NilHandler(thisPaperSource["authors"].([]interface{})[0].(map[string]interface{})["name"], "string").(string) +
			strconv.Itoa(NilHandler(thisPaperSource["year"], "int").(int)),
		Size: NilHandler(thisPaperSource["n_citation"], "int").(int),
		Type: "major",
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

	references := NilHandler(thisPaperSource["relateds"], "list").([]interface{})
	referenceIds := make([]string, 0)
	for _, reference := range references {
		referenceIds = append(referenceIds, reference.(string))
	}

	DFSSimilar(referenceIds, majorNode, 0)

	for _, node := range nodes {
		node.Size = GetSize(node.Size, maxCitation, minCitation)
		node.Style.Fill = GetColor(GetD(node.Info.Year, maxYear, minYear))
	}

	resp = &types.PaperSimilarNetResponse{
		Nodes: nodes,
		Edges: edges,
	}
	return resp, nil
}

func DFSSimilar(referenceIds []string, fatherNode types.PaperNodeJSON, level int) {
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
		if paper.(map[string]interface{})["found"].(bool) == false {
			continue
		}
		source := paper.(map[string]interface{})["_source"].(map[string]interface{})
		node := types.PaperNodeJSON{
			Id:    NilHandler(source["id"], "string").(string),
			Label: NilHandler(source["authors"].([]interface{})[0].(map[string]interface{})["name"], "string").(string) + strconv.Itoa(NilHandler(source["year"], "int").(int)),
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

		references := NilHandler(source["relateds"], "list").([]interface{})
		referenceIds = make([]string, 0)
		for _, reference := range references {
			referenceIds = append(referenceIds, reference.(string))
		}

		DFSSimilar(referenceIds, node, level+1)
	}
}
