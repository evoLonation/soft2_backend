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
			Name:  NilHandler(author.(map[string]interface{})["name"], "string").(string),
			Id:    NilHandler(author.(map[string]interface{})["id"], "string").(string),
			HasId: hasId,
		})
	}
	var urlStrings []string
	urls := NilHandler(source["url"], "list").([]interface{})
	for _, url := range urls {
		urlStrings = append(urlStrings, NilHandler(url, "string").(string))
	}

	var keywordStrings []string
	keywords := NilHandler(source["keywords"], "list").([]interface{})
	for _, keyword := range keywords {
		keywordStrings = append(keywordStrings, keyword.(string))
	}

	var referencePapers []types.PaperJSON
	var referenceIds []string
	references := NilHandler(source["references"], "list").([]interface{})
	for _, reference := range references {
		referenceIds = append(referenceIds, reference.(string))
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
		authors := NilHandler(source["authors"].([]interface{}), "list").([]interface{})
		referencePapers = append(referencePapers, types.PaperJSON{
			Id:     NilHandler(source["id"], "string").(string),
			Title:  NilHandler(source["title"], "string").(string),
			Author: NilHandler(authors[0].(map[string]interface{})["name"], "string").(string),
			Year:   NilHandler(source["year"], "int").(int),
		})
	}

	var SimilarPapers []types.PaperJSON
	var similarIds []string
	similar := NilHandler(source["relateds"], "list").([]interface{})
	for _, sim := range similar {
		similarIds = append(similarIds, sim.(string))
	}
	var similarBuf bytes.Buffer
	similarQuery := map[string]interface{}{
		"ids": similarIds,
	}
	if err := json.NewEncoder(&similarBuf).Encode(similarQuery); err != nil {
		log.Printf("encode query error: %v", err)
	}
	similarRes := database.MgetPaper(similarBuf)
	papers = NilHandler(similarRes["docs"], "list").([]interface{})
	for _, paper := range papers {
		source := paper.(map[string]interface{})["_source"].(map[string]interface{})
		authors := NilHandler(source["authors"].([]interface{}), "list").([]interface{})
		SimilarPapers = append(SimilarPapers, types.PaperJSON{
			Id:     NilHandler(source["id"], "string").(string),
			Title:  NilHandler(source["title"], "string").(string),
			Author: NilHandler(authors[0].(map[string]interface{})["name"], "string").(string),
			Year:   NilHandler(source["year"], "int").(int),
		})
	}

	resp = &types.PaperDetailResponse{
		Id:         NilHandler(source["id"], "string").(string),
		Title:      NilHandler(source["title"], "string").(string),
		Abstract:   NilHandler(source["abstract"], "string").(string),
		Authors:    authors,
		Doi:        NilHandler(source["doi"], "string").(string),
		ISBN:       NilHandler(source["isbn"], "string").(string),
		Org:        NilHandler(source["org"], "string").(string),
		Keywords:   keywordStrings,
		Year:       NilHandler(source["year"], "int").(int),
		NCitation:  NilHandler(source["n_citation"], "int").(int),
		Publisher:  NilHandler(source["publisher"], "string").(string),
		References: referencePapers,
		Similars:   SimilarPapers,
		Urls:       urlStrings,
	}
	return resp, nil
}
