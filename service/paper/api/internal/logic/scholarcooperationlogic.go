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

type ScholarCooperationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScholarCooperationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScholarCooperationLogic {
	return &ScholarCooperationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScholarCooperationLogic) ScholarCooperation(req *types.ScholarCooperationRequest) (resp *types.ScholarCooperationResponse, err error) {
	// todo: add your logic here and delete this line
	var authorBuf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"id": req.ScholarId,
			},
		},
	}
	if err := json.NewEncoder(&authorBuf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s\n", err)
	}
	log.Println(authorBuf.String())
	res := database.SearchAuthor(authorBuf)

	source := res["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})
	pubs := source["pubs"].([]interface{})
	var pubIds []string
	for _, pub := range pubs {
		pubIds = append(pubIds, pub.(map[string]interface{})["i"].(string))
	}
	var paperBuf bytes.Buffer
	mget := map[string]interface{}{
		"ids": pubIds,
	}
	if err := json.NewEncoder(&paperBuf).Encode(mget); err != nil {
		log.Printf("Error encoding query: %s\n", err)
	}
	log.Println(paperBuf.String())
	res = database.MgetPaper(paperBuf)

	papers := res["docs"].([]interface{})
	coopList := make(map[string]types.CoopJSON)
	for _, paper := range papers {
		authors := paper.(map[string]interface{})["_source"].(map[string]interface{})["authors"].([]interface{})
		for _, author := range authors {
			v, ok := coopList[author.(map[string]interface{})["id"].(string)]
			if ok {
				v.Time++
				coopList[author.(map[string]interface{})["id"].(string)] = v
			} else {
				coopJSON := types.CoopJSON{
					ScholarId: author.(map[string]interface{})["id"].(string),
					Name:      author.(map[string]interface{})["name"].(string),
					Time:      1,
				}
				var coopBuf bytes.Buffer
				authorQuery := map[string]interface{}{
					"query": map[string]interface{}{
						"match": map[string]interface{}{
							"id": author.(map[string]interface{})["id"].(string),
						},
					},
				}
				if err := json.NewEncoder(&coopBuf).Encode(authorQuery); err != nil {
					log.Printf("Error encoding query: %s\n", err)
				}
				log.Println(coopBuf.String())
				res = database.SearchAuthor(coopBuf)
				source := res["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})
				coopJSON.Institution = source["orgs"].([]interface{})[0].(string)
				coopList[author.(map[string]interface{})["id"].(string)] = coopJSON
			}
		}
	}

	var coopJSONs []types.CoopJSON
	for _, v := range coopList {
		coopJSONs = append(coopJSONs, v)
	}
	resp = &types.ScholarCooperationResponse{
		CoopList: coopJSONs,
	}
	return resp, nil
}
