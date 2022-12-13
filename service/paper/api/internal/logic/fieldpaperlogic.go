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

type FieldPaperLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFieldPaperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FieldPaperLogic {
	return &FieldPaperLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FieldPaperLogic) FieldPaper(req *types.FieldPaperRequest) (resp *types.FieldPaperResponse, err error) {
	// todo: add your logic here and delete this line
	var paperBuf bytes.Buffer
	paperQuery := map[string]interface{}{
		"from": req.Start,
		"size": req.End - req.Start,
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"keywords": req.Field,
			},
		},
		"sort": map[string]interface{}{
			"n_citation": map[string]interface{}{
				"order": "desc",
			},
		},
	}
	if err := json.NewEncoder(&paperBuf).Encode(paperQuery); err != nil {
		log.Printf("Error encoding query: %s\n", err)
	}
	log.Println(paperBuf.String())
	paperRes := database.SearchPaper(paperBuf)
	log.Printf("finished searching paper in field %v", req.Field)

	var papers []types.PaperResponseJSON
	hits := paperRes["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		var authorJSONs []types.AuthorJSON
		authors := NilHandler(source["authors"], "list").([]interface{})
		for _, author := range authors {
			hasId := false
			if author.(map[string]interface{})["id"] != nil {
				hasId = true
			}
			authorJSONs = append(authorJSONs, types.AuthorJSON{
				Name:  NilHandler(author.(map[string]interface{})["name"], "string").(string),
				Id:    NilHandler(author.(map[string]interface{})["id"], "string").(string),
				HasId: hasId,
			})
		}
		papers = append(papers, types.PaperResponseJSON{
			Id:        NilHandler(source["id"], "string").(string),
			Title:     NilHandler(source["title"], "string").(string),
			Abstract:  NilHandler(source["abstract"], "string").(string),
			Authors:   authorJSONs,
			Year:      NilHandler(source["year"], "int").(int),
			NCitation: NilHandler(source["n_citation"], "int").(int),
			Publisher: NilHandler(source["venue"], "string").(string),
		})
	}

	paperNum := int(paperRes["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))
	if paperNum > 10000 {
		paperNum = 10000
	}
	resp = &types.FieldPaperResponse{
		PaperNum: paperNum,
		Papers:   papers,
	}
	return resp, nil
}
