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

type PaperLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaperLogic {
	return &PaperLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaperLogic) Paper(req *types.PaperRequest) (resp *types.PaperResponse, err error) {
	// todo: add your logic here and delete this line
	searchContent := req.Content
	var buf bytes.Buffer
	var must []map[string]interface{}
	var should []map[string]interface{}
	var mustNot []map[string]interface{}
	query := map[string]interface{}{
		"from": req.Start,
		"size": req.End - req.Start,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": map[string]interface{}{
					"range": map[string]interface{}{
						"year": map[string]interface{}{
							"gte": req.StartYear,
							"lte": req.EndYear,
						},
					},
				},
			},
		},
	}
	for _, content := range searchContent {
		if content.Type == 0 { // 0 -> and -> must
			if content.IsExact == 0 {
				must = append(must, map[string]interface{}{
					"match_phrase": map[string]interface{}{
						TranslateSearchKey(content.SearchType): content.Content,
					},
				})
			} else {
				must = append(must, map[string]interface{}{
					"match": map[string]interface{}{
						TranslateSearchKey(content.SearchType): content.Content,
					},
				})
			}
		} else if content.Type == 1 { // 1 -> or -> should
			if content.IsExact == 0 {
				should = append(should, map[string]interface{}{
					"match_phrase": map[string]interface{}{
						TranslateSearchKey(content.SearchType): content.Content,
					},
				})
			} else {
				should = append(should, map[string]interface{}{
					"match": map[string]interface{}{
						TranslateSearchKey(content.SearchType): content.Content,
					},
				})
			}
		} else if content.Type == 2 { // 2 -> not -> must_not
			if content.IsExact == 0 {
				mustNot = append(mustNot, map[string]interface{}{
					"match_phrase": map[string]interface{}{
						TranslateSearchKey(content.SearchType): content.Content,
					},
				})
			} else {
				mustNot = append(mustNot, map[string]interface{}{
					"match": map[string]interface{}{
						TranslateSearchKey(content.SearchType): content.Content,
					},
				})
			}
		}
	}
	if len(must) != 0 {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = must
	}
	if len(should) != 0 {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["should"] = should
	}
	if len(mustNot) != 0 {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must_not"] = mustNot
	}
	if req.SortType == 1 {
		query["sort"] = map[string]interface{}{
			"n_citation": map[string]interface{}{
				"order": "desc",
			},
		}
	} else if req.SortType == 2 {
		query["sort"] = map[string]interface{}{
			"year": map[string]interface{}{
				"order": "desc",
			},
		}
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s\n", err)
	}
	log.Println(buf.String())
	res := database.SearchPaper(buf)

	var papers []types.PaperResponseJSON
	var themes map[string]int
	var years []int
	for _, hit := range res["hits"].(map[string]interface{})["hits"].([]interface{}) {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
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
		papers = append(papers, types.PaperResponseJSON{
			Id:        NilHandler(source["id"], "string").(string),
			Title:     NilHandler(source["title"], "string").(string),
			Abstract:  NilHandler(source["abstract"], "string").(string),
			Authors:   authors,
			Year:      NilHandler(source["year"], "int").(int),
			NCitation: NilHandler(source["n_citation"], "int").(int),
			Publisher: NilHandler(source["venue"], "string").(string),
		})
		log.Println(source["keywords"])
		keywords := NilHandler(source["keywords"], "list").([]interface{})
		cnt := 0
		themes = make(map[string]int)
		for _, theme := range keywords {
			themes[theme.(string)] = cnt
			cnt++
		}
		years = append(years, int(source["year"].(float64)))
	}

	paperNum := int(res["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))
	if paperNum > 10000 {
		paperNum = 10000
	}
	resp = &types.PaperResponse{
		PaperNum: paperNum,
		Papers:   papers,
		Themes:   getKeywords(themes),
		Years:    years,
	}
	return resp, nil
}

func TranslateSearchKey(searchKey int) string {
	switch searchKey {
	case 0:
		return "title"
	case 1:
		return "authors.name"
	case 2:
		return "keywords"
	case 3:
		return "abstract"
	case 4:
		return "doi"
	case 5:
		return "venue"
	case 6:
		return "author.org"
	case 7:
		return "year"
	default:
		return ""
	}
}

func getKeywords(source map[string]int) []string {
	var keywords []string
	for k := range source {
		keywords = append(keywords, k)
	}
	return keywords
}
