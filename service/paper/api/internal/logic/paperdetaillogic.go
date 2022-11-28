package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"soft2_backend/service/paper/api/internal/svc"
	"soft2_backend/service/paper/api/internal/types"
	"soft2_backend/service/paper/database"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaperDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaperDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaperDetailLogic {
	return &PaperDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaperDetailLogic) PaperDetail(req *types.PaperDetailRequest) (resp *types.PaperDetailResponse, err error) {
	// todo: add your logic here and delete this line
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"id": req.Id,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s\n", err)
	}
	log.Println(buf.String())
	res := database.SearchPaper(buf)

	if res["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64) == 0 {
		return nil, nil
	}
	source := res["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})
	var authors []types.AuthorJSON
	for _, author := range source["authors"].([]interface{}) {
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
	var urls []string
	for _, url := range source["url"].([]interface{}) {
		urls = append(urls, NilHandler(url, "string").(string))
	}
	resp = &types.PaperDetailResponse{
		Title:     source["title"].(string),
		Abstract:  NilHandler(source["abstract"], "string").(string),
		Authors:   authors,
		Doi:       NilHandler(source["doi"], "string").(string),
		ISBN:      NilHandler(source["isbn"], "string").(string),
		Org:       NilHandler(source["org"], "string").(string),
		Keywords:  NilHandler(source["keywords"], "list").([]string),
		Year:      NilHandler(source["year"], "int").(int),
		NCitation: NilHandler(source["n_citation"], "int").(int),
		Publisher: NilHandler(source["publisher"], "string").(string),
		Urls:      urls,
	}
	return resp, nil
}
