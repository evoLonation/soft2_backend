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

type FieldScholarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFieldScholarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FieldScholarLogic {
	return &FieldScholarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FieldScholarLogic) FieldScholar(req *types.FieldScholarRequest) (resp *types.FieldScholarResponse, err error) {
	// todo: add your logic here and delete this line
	var scholarBuf bytes.Buffer
	scholarQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"tags.t": req.Field,
			},
		},
	}
	if err := json.NewEncoder(&scholarBuf).Encode(scholarQuery); err != nil {
		log.Printf("Error encoding query: %s", err)
	}
	log.Println(scholarBuf.String())
	scholarRes := database.SearchAuthor(scholarBuf)

	var scholars []types.FieldScholarJSON
	hits := scholarRes["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		tags := NilHandler(source["tags"], "list").([]interface{})
		var totalWeight float64
		var thisWeight float64
		var minLevenshtein = 100.0
		for _, tag := range tags {
			totalWeight += float64(NilHandler(tag.(map[string]interface{})["w"], "int").(int))
			if Levenshtein(req.Field, tag.(map[string]interface{})["t"].(string), 1, 1, 1) < minLevenshtein {
				minLevenshtein = Levenshtein(req.Field, tag.(map[string]interface{})["t"].(string), 1, 1, 1)
				thisWeight = tag.(map[string]interface{})["w"].(float64)
			}
		}

		scholars = append(scholars, types.FieldScholarJSON{
			ScholarId: NilHandler(source["id"], "string").(string),
			Name:      NilHandler(source["name"], "string").(string),
			NPaper:    NilHandler(source["n_pubs"], "int").(int),
			NCitation: NilHandler(source["n_citation"], "int").(int),
			Weight:    thisWeight / totalWeight,
		})
	}

	scholarNum := int(scholarRes["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))
	if scholarNum > 10000 {
		scholarNum = 10000
	}
	resp = &types.FieldScholarResponse{
		ScholarNum: scholarNum,
		Scholars:   scholars,
	}
	return resp, nil
}
