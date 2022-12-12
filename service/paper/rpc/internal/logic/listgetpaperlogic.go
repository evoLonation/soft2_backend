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

type ListGetPaperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListGetPaperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListGetPaperLogic {
	return &ListGetPaperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListGetPaperLogic) ListGetPaper(in *paper.ListGetPaperReq) (*paper.ListGetPaperReply, error) {
	// todo: add your logic here and delete this line
	var buf bytes.Buffer
	var idList []string
	for _, id := range in.PaperId {
		idList = append(idList, id)
	}
	query := map[string]interface{}{
		"ids": idList,
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s\n", err)
	}
	log.Println(buf.String())
	res := database.MgetPaper(buf)
	log.Println(res)

	papers := NilHandler(res["docs"], "list").([]interface{})
	paperList := make([]*paper.GetPaperReply, 0)
	for _, p := range papers {
		if p.(map[string]interface{})["found"].(bool) == false {
			continue
		}
		source := p.(map[string]interface{})["_source"].(map[string]interface{})
		var authors []*paper.AuthorJSON
		for _, author := range source["authors"].([]interface{}) {
			hasId := false
			if author.(map[string]interface{})["id"] != nil {
				hasId = true
			}
			authors = append(authors, &paper.AuthorJSON{
				Name:  NilHandler(author.(map[string]interface{})["name"], "string").(string),
				Id:    NilHandler(author.(map[string]interface{})["id"], "string").(string),
				HasId: hasId,
			})
		}

		firstAuthorOrg := ""
		if len(authors) > 0 {
			var authorBuf bytes.Buffer
			firstAuthorId := authors[0].Id
			firstAuthorQuery := map[string]interface{}{
				"query": map[string]interface{}{
					"match": map[string]interface{}{
						"id": firstAuthorId,
					},
				},
			}
			if err := json.NewEncoder(&authorBuf).Encode(firstAuthorQuery); err != nil {
				log.Printf("Error encoding query: %s\n", err)
			}
			log.Println(authorBuf.String())
			firstAuthorRes := database.SearchAuthor(authorBuf)
			firstAuthorHits := NilHandler(firstAuthorRes["hits"].(map[string]interface{})["hits"], "list").([]interface{})
			firstAuthorSource := firstAuthorHits[0].(map[string]interface{})["_source"].(map[string]interface{})
			firstAuthorOrg = NilHandler(firstAuthorSource["orgs"].([]interface{})[0], "string").(string)
		}

		paperList = append(paperList, &paper.GetPaperReply{
			PaperName: source["title"].(string),
			Authors:   authors,
			Org:       firstAuthorOrg,
			Year:      int64(NilHandler(source["year"], "int").(int)),
		})
	}

	return &paper.ListGetPaperReply{
		Papers: paperList,
	}, nil
}
