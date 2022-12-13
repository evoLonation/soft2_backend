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
		{"computer science", "machine learning", "computer vision"},
		{"mathematics", "linear algebra", "calculus"},
		{"physics", "quantum mechanics", "electromagnetism", "chemistry"},
		{"biology", "genetics", "ecology", "zoology"},
		{"economics", "microeconomics", "macroeconomics"},
		{"psychology", "cognitive psychology", "social psychology"},
		{"history", "ancient history", "modern history"},
		{"environment", "climate change", "global warming"}}
	areaJsonList := make([]types.AreaJSON, 0)
	for _, area := range areas {
		var paperBuf bytes.Buffer
		paperQuery := map[string]interface{}{
			"from": 0,
			"size": req.PaperNum,
			"query": map[string]interface{}{
				"bool": map[string]interface{}{},
			},
			"sort": map[string]interface{}{
				"n_citation": map[string]interface{}{
					"order": "desc",
				},
			},
		}
		var paperShould []map[string]interface{}
		for _, subArea := range area {
			paperShould = append(paperShould, map[string]interface{}{
				"match_phrase": map[string]interface{}{
					"keywords": subArea,
				},
			})
		}
		paperQuery["query"].(map[string]interface{})["bool"].(map[string]interface{})["should"] = paperShould
		if err := json.NewEncoder(&paperBuf).Encode(paperQuery); err != nil {
			log.Printf("Error encoding query: %s", err)
		}
		log.Println(paperBuf.String())
		paperResult := database.SearchPaper(paperBuf)

		var paperList []types.PaperInfoJSON
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

		var scholarBuf bytes.Buffer
		scholarQuery := map[string]interface{}{
			"from": 0,
			"size": req.ScholarNum,
			"query": map[string]interface{}{
				"bool": map[string]interface{}{},
			},
			"sort": map[string]interface{}{
				"n_citation": map[string]interface{}{
					"order": "desc",
				},
			},
		}
		var scholarShould []map[string]interface{}
		for _, subArea := range area {
			scholarShould = append(scholarShould, map[string]interface{}{
				"match_phrase": map[string]interface{}{
					"tags.t": subArea,
				},
			})
		}
		scholarQuery["query"].(map[string]interface{})["bool"].(map[string]interface{})["should"] = scholarShould
		if err := json.NewEncoder(&scholarBuf).Encode(scholarQuery); err != nil {
			log.Printf("Error encoding query: %s", err)
		}
		log.Println(scholarBuf.String())
		scholarResult := database.SearchAuthor(scholarBuf)

		var scholarList []types.ScholarInfoJSON
		hits = scholarResult["hits"].(map[string]interface{})["hits"].([]interface{})
		for _, hit := range hits {
			source := hit.(map[string]interface{})["_source"].(map[string]interface{})
			scholarList = append(scholarList, types.ScholarInfoJSON{
				ScholarId: NilHandler(source["id"], "string").(string),
				Name:      NilHandler(source["name"], "string").(string),
				RefNum:    NilHandler(source["n_citation"], "int").(int),
			})
		}

		areaJsonList = append(areaJsonList, types.AreaJSON{
			Type:     area,
			Papers:   paperList,
			Scholars: scholarList,
		})
	}
	resp = &types.HomeInfoResponse{
		Areas: areaJsonList,
	}
	return resp, nil
}
