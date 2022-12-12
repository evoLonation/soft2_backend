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

type ScholarBarchartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScholarBarchartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScholarBarchartLogic {
	return &ScholarBarchartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScholarBarchartLogic) ScholarBarchart(req *types.ScholarBarchartRequest) (resp *types.ScholarBarchartResponse, err error) {
	// todo: add your logic here and delete this line
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"id": req.ScholarId,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s\n", err)
	}
	log.Println(buf.String())
	res := database.SearchAuthor(buf)
	log.Println(res)

	var statistic = make([]types.StatisticJSON, 0)
	source := res["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})
	log.Println(source)
	statistics := NilHandler(source["statistics"], "list").([]interface{})
	log.Println(statistics)
	for _, s := range statistics {
		s = append(statistic, types.StatisticJSON{
			Year:         NilHandler(s.(map[string]interface{})["year"], "int").(int),
			Achievements: NilHandler(s.(map[string]interface{})["n_pubs"], "int").(int),
			References:   NilHandler(s.(map[string]interface{})["n_citation"], "int").(int),
		})
	}

	//var statistic = make([]types.StatisticJSON, 0)
	//var achievementsMap = make(map[int]int)
	//var referencesMap = make(map[int]int)
	//source := res["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})
	//pubs := NilHandler(source["pubs"], "list").([]interface{})
	//for _, pub := range pubs {
	//	var paperBuf bytes.Buffer
	//	paperQuery := map[string]interface{}{
	//		"query": map[string]interface{}{
	//			"match": map[string]interface{}{
	//				"id": pub.(map[string]interface{})["i"].(string),
	//			},
	//		},
	//	}
	//	if err := json.NewEncoder(&paperBuf).Encode(paperQuery); err != nil {
	//		log.Printf("Error encoding query: %s\n", err)
	//	}
	//	log.Println(paperBuf.String())
	//	paperRes := database.SearchPaper(paperBuf)
	//
	//	paperSource := paperRes["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})
	//	year := NilHandler(paperSource["year"], "int").(int)
	//	achievementsMap[year]++
	//	referencesMap[year] += NilHandler(paperSource["n_citation"], "int").(int)
	//}
	//
	//for k, v := range achievementsMap {
	//	statistic = append(statistic, types.StatisticJSON{
	//		Year:         k,
	//		Achievements: v,
	//		References:   referencesMap[k],
	//	})
	//}

	resp = &types.ScholarBarchartResponse{
		Statistic: statistic,
	}
	return
}
