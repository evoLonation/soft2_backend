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

type ScholarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScholarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScholarLogic {
	return &ScholarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScholarLogic) Scholar(req *types.ScholarRequest) (resp *types.ScholarResponse, err error) {
	// todo: add your logic here and delete this line
	var buf bytes.Buffer
	query := map[string]interface{}{
		"from": req.Start,
		"size": req.End - req.Start,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []map[string]interface{}{
					{
						"match": map[string]interface{}{
							"name": req.Name,
						},
					},
					{
						"match": map[string]interface{}{
							"orgs": req.Institution,
						},
					},
				},
			},
		},
		"sort": map[string]interface{}{
			"_score": map[string]interface{}{
				"order": "desc",
			},
			"n_citation": map[string]interface{}{
				"order": "desc",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s\n", err)
	}
	log.Println(buf.String())
	res := database.SearchAuthor(buf)

	var scholars []types.ScholarResponseJSON
	for _, hit := range res["hits"].(map[string]interface{})["hits"].([]interface{}) {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		orgs := NilHandler(source["orgs"], "list").([]interface{})
		var org string
		if len(orgs) == 0 {
			org = ""
		} else {
			org = orgs[0].(string)
		}
		scholar := types.ScholarResponseJSON{
			Id:          NilHandler(source["id"], "string").(string),
			Name:        NilHandler(source["name"], "string").(string),
			Institution: org,
			PaperNum:    NilHandler(source["n_pubs"], "int").(int),
		}
		scholars = append(scholars, scholar)
	}

	scholarNum := int(res["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))
	if scholarNum > 10000 {
		scholarNum = 10000
	}
	resp = &types.ScholarResponse{
		ScholarNum: scholarNum,
		Scholar:    scholars,
	}
	return
}
