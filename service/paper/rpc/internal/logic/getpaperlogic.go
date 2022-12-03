package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"soft2_backend/service/paper/database"

	"soft2_backend/service/paper/rpc/internal/svc"
	"soft2_backend/service/paper/rpc/paper"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPaperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPaperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPaperLogic {
	return &GetPaperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPaperLogic) GetPaper(in *paper.GetPaperReq) (*paper.GetPaperReply, error) {
	// todo: add your logic here and delete this line
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"id": in.PaperId,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s\n", err)
	}
	log.Println(buf.String())
	res := database.SearchPaper(buf)

	source := res["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})
	var authors []*paper.AuthorJSON
	for _, author := range source["authors"].([]interface{}) {
		hasId := false
		if author.(map[string]interface{})["id"] != nil {
			hasId = true
		}
		authors = append(authors, &paper.AuthorJSON{
			Name:  author.(map[string]interface{})["name"].(string),
			Id:    NilHandler(author.(map[string]interface{})["id"], "string").(string),
			HasId: hasId,
		})
	}

	firstAuthorId := authors[0].Id
	firstAuthorQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"id": firstAuthorId,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(firstAuthorQuery); err != nil {
		log.Printf("Error encoding query: %s\n", err)
	}
	log.Println(buf.String())
	firstAuthorRes := database.SearchAuthor(buf)

	firstAuthorSource := firstAuthorRes["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})
	firstAuthorOrg := firstAuthorSource["orgs"].([]interface{})[0].(string)

	return &paper.GetPaperReply{
		PaperName: source["title"].(string),
		Authors:   authors,
		Org:       firstAuthorOrg,
		Year:      NilHandler(source["year"], "int").(int64),
	}, nil
}
