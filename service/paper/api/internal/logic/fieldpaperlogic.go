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
	}
	if err := json.NewEncoder(&paperBuf).Encode(paperQuery); err != nil {
		log.Printf("Error encoding query: %s\n", err)
	}
	log.Println(paperBuf.String())
	paperRes := database.SearchPaper(paperBuf)

	var papers []types.PaperResponseJSON
	hits := paperRes["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		var authors []types.AuthorJSON
		for _, author := range source["authors"].([]interface{}) {
			hasId := false
			if author.(map[string]interface{})["id"] != nil {
				hasId = true
			}
			authors = append(authors, types.AuthorJSON{
				Name:  NilHandler(author.(map[string]interface{})["name"], "string").(string),
				Id:    NilHandler(author.(map[string]interface{})["id"], "string").(string),
				HasId: hasId,
			})
		}
		papers = append(papers, types.PaperResponseJSON{
			Title:     NilHandler(source["title"], "string").(string),
			Abstract:  NilHandler(source["abstract"], "string").(string),
			Authors:   authors,
			Year:      NilHandler(source["year"], "int").(int),
			NCitation: NilHandler(source["n_citation"], "int").(int),
			Publisher: NilHandler(source["venue"], "string").(string),
		})
	}

	resp = &types.FieldPaperResponse{
		PaperNum: int(paperRes["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		Papers:   papers,
	}
	return resp, nil
}
