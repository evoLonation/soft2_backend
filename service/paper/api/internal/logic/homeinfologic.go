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

type HomeInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeInfoLogic {
	return &HomeInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeInfoLogic) HomeInfo(req *types.HomeInfoRequest) (resp *types.HomeInfoResponse, err error) {
	// todo: add your logic here and delete this line
	areas := [][]string{
		{"deep learning", "machine learning", "artificial intelligence"},
		{"mathematics", "linear algebra", "calculus"},
		{"physics", "quantum mechanics", "electromagnetism", "chemistry"},
		{"biology", "genetics", "ecology", "zoology"},
		{"economics", "microeconomics", "macroeconomics"},
		{"psychology", "cognitive psychology", "social psychology"},
		{"history", "ancient history", "modern history"},
		{"environment", "climate change", "global warming"}}
	areasNum := req.AreasNum
	if areasNum == 0 {
		areasNum = len(areas)
	}
	areaJsonList := make([]types.AreaJSON, 0)
	for i, area := range areas {
		if i == areasNum {
			break
		}
		paperQueryString, scholarQueryString := GenerateQueryString(area)
		var paperList []types.PaperInfoJSON
		var scholarList []types.ScholarInfoJSON
		go func() {
			var paperBuf bytes.Buffer
			paperQuery := map[string]interface{}{
				"from": 0,
				"size": req.PaperNum,
				"query": map[string]interface{}{
					"query_string": map[string]interface{}{
						"query": paperQueryString,
					},
				},
				//"aggs": map[string]interface{}{
				//	"journals": map[string]interface{}{
				//		"terms": map[string]interface{}{
				//			"field": "venue.filter",
				//			"order": map[string]interface{}{
				//				"_count": "desc",
				//			},
				//		},
				//	},
				//},
			}
			if err := json.NewEncoder(&paperBuf).Encode(paperQuery); err != nil {
				log.Printf("Error encoding query: %s", err)
			}
			log.Println(paperBuf.String())
			paperResult := database.SearchPaper(paperBuf)
			log.Println(paperResult)

			hits := paperResult["hits"].(map[string]interface{})["hits"].([]interface{})
			for _, hit := range hits {
				var authorList []string
				source := hit.(map[string]interface{})["_source"].(map[string]interface{})
				authors := NilHandler(source["authors"], "list").([]interface{})
				for _, author := range authors {
					authorList = append(authorList, NilHandler(author.(map[string]interface{})["name"], "string").(string))
				}
				paperList = append(paperList, types.PaperInfoJSON{
					Title:     NilHandler(source["title"], "string").(string),
					Authors:   authorList,
					NCitation: NilHandler(source["n_citation"], "int").(int),
				})
			}
		}()

		go func() {
			var scholarBuf bytes.Buffer
			scholarQuery := map[string]interface{}{
				"from": 0,
				"size": req.PaperNum,
				"query": map[string]interface{}{
					"query_string": map[string]interface{}{
						"query": scholarQueryString,
					},
				},
			}
			if err := json.NewEncoder(&scholarBuf).Encode(scholarQuery); err != nil {
				log.Printf("Error encoding query: %s", err)
			}
			log.Println(scholarBuf.String())
			scholarResult := database.SearchAuthor(scholarBuf)

			hits := scholarResult["hits"].(map[string]interface{})["hits"].([]interface{})
			for _, hit := range hits {
				source := hit.(map[string]interface{})["_source"].(map[string]interface{})
				scholarList = append(scholarList, types.ScholarInfoJSON{
					ScholarId: NilHandler(source["id"], "string").(string),
					Name:      NilHandler(source["name"], "string").(string),
					RefNum:    NilHandler(source["n_citation"], "int").(int),
				})
			}
		}()

		areaJsonList = append(areaJsonList, types.AreaJSON{
			Type:     areas[0],
			Papers:   paperList,
			Scholars: scholarList,
		})
	}

	resp = &types.HomeInfoResponse{
		Areas: areaJsonList,
	}
	return resp, nil
}

func GenerateQueryString(areas []string) (string, string) {
	paperQueryString := ""
	scholarQueryString := ""
	for i, area := range areas {
		if i == 0 {
			paperQueryString += "keywords:\"" + area + "\""
			scholarQueryString += "tags.t:\"" + area + "\""
		} else {
			paperQueryString += " OR keywords:\"" + area + "\""
			scholarQueryString += " OR tags.t:\"" + area + "\""
		}
	}
	return paperQueryString, scholarQueryString
}
