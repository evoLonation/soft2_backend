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
		{"computer science", "machine learning", "artificial intelligence"},
		{"mathematics", "geometry", "calculus"},
		{"physics", "electricity", "optics"},
		{"biology", "genetics", "ecology"},
		{"economics", "microeconomics", "macroeconomics"},
		{"psychology", "cognitive", "social"},
		{"environment", "environment", "global warming"}}
	areasNum := req.AreasNum
	if areasNum == 0 {
		areasNum = len(areas)
	}

	paperChan := make(chan types.PaperInfoJSON, areasNum*req.PaperNum)
	scholarChan := make(chan types.ScholarInfoJSON, areasNum*req.ScholarNum)
	areaChan := make(chan types.AreaJSON, 7)
	areaJsonList := make([]types.AreaJSON, 0)
	for i, _ := range areas {
		if i == areasNum {
			break
		}
		area := areas[i]
		go func() {
			paperQueryString, scholarQueryString := GenerateQueryString(area[1:])
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
				log.Printf("finished searching paper in area: %v", area)

				hits := paperResult["hits"].(map[string]interface{})["hits"].([]interface{})
				for _, hit := range hits {
					var authorList []string
					source := hit.(map[string]interface{})["_source"].(map[string]interface{})
					authors := NilHandler(source["authors"], "list").([]interface{})
					for _, author := range authors {
						authorList = append(authorList, NilHandler(author.(map[string]interface{})["name"], "string").(string))
					}
					thisPaperJson := types.PaperInfoJSON{
						Title:     NilHandler(source["title"], "string").(string),
						Authors:   authorList,
						NCitation: NilHandler(source["n_citation"], "int").(int),
					}
					paperChan <- thisPaperJson
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
				log.Printf("finished searching scholar in area: %v", area)

				hits := scholarResult["hits"].(map[string]interface{})["hits"].([]interface{})
				for _, hit := range hits {
					source := hit.(map[string]interface{})["_source"].(map[string]interface{})
					thisScholarJson := types.ScholarInfoJSON{
						ScholarId: NilHandler(source["id"], "string").(string),
						Name:      NilHandler(source["name"], "string").(string),
						RefNum:    NilHandler(source["n_citation"], "int").(int),
					}
					scholarChan <- thisScholarJson
				}
			}()

			for j := 0; j < req.PaperNum; j++ {
				paperList = append(paperList, <-paperChan)
			}
			for j := 0; j < req.ScholarNum; j++ {
				scholarList = append(scholarList, <-scholarChan)
			}
			thisAreaJson := types.AreaJSON{
				Type:     area,
				Papers:   paperList,
				Scholars: scholarList,
			}
			areaChan <- thisAreaJson
		}()
	}
	for i := 0; i < areasNum; i++ {
		areaJsonList = append(areaJsonList, <-areaChan)
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
