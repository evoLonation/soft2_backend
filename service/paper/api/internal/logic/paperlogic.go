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
		"from": (req.Page - 1) * 10,
		"size": 10,
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
		"_source": []string{
			"title",
			"abstract",
			"authors",
			"year",
			"n_citation",
			"venue",
			"keywords",
		},
	}
	for _, content := range searchContent {
		if content.Type == 0 { // 0 -> and -> must
			if content.IsExact == 0 {
				must = append(must, map[string]interface{}{
					"match_phrase": map[string]interface{}{
						TranslateSearchKey(int(content.SearchType)): content.Content,
					},
				})
			} else {
				must = append(must, map[string]interface{}{
					"match": map[string]interface{}{
						TranslateSearchKey(int(content.SearchType)): content.Content,
					},
				})
			}
		} else if content.Type == 1 { // 1 -> or -> should
			if content.IsExact == 0 {
				should = append(should, map[string]interface{}{
					"match_phrase": map[string]interface{}{
						TranslateSearchKey(int(content.SearchType)): content.Content,
					},
				})
			} else {
				should = append(should, map[string]interface{}{
					"match": map[string]interface{}{
						TranslateSearchKey(int(content.SearchType)): content.Content,
					},
				})
			}
		} else if content.Type == 2 { // 2 -> not -> must_not
			if content.IsExact == 0 {
				mustNot = append(mustNot, map[string]interface{}{
					"match_phrase": map[string]interface{}{
						TranslateSearchKey(int(content.SearchType)): content.Content,
					},
				})
			} else {
				mustNot = append(mustNot, map[string]interface{}{
					"match": map[string]interface{}{
						TranslateSearchKey(int(content.SearchType)): content.Content,
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
	var themes []string
	var years []int
	for _, hit := range res["hits"].(map[string]interface{})["hits"].([]interface{}) {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		var authors []types.AuthorJSON
		for _, author := range source["authors"].([]interface{}) {
			authors = append(authors, types.AuthorJSON{
				Name: author.(map[string]interface{})["name"].(string),
				Id:   NilHandler(author.(map[string]interface{})["id"], "string").(string),
			})
		}
		papers = append(papers, types.PaperResponseJSON{
			Title:     source["title"].(string),
			Abstract:  NilHandler(source["abstract"], "string").(string),
			Authors:   authors,
			Year:      NilHandler(source["year"], "int").(int),
			NCitation: NilHandler(source["n_citation"], "int").(int),
			Publisher: NilHandler(source["venue"].(map[string]interface{})["raw"], "string").(string),
		})
		log.Println(papers)

		keywords := NilHandler(source["keywords"], "list").([]interface{})
		for _, theme := range keywords {
			themes = append(themes, theme.(string))
		}
		years = append(years, int(source["year"].(float64)))
	}

	resp = &types.PaperResponse{
		PageNum:  int(res["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))/10 + 1,
		PaperNum: int(res["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		Papers:   papers,
		Themes:   themes,
		Years:    years,
		Fields:   themes,
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
		return "venue.raw"
	case 6:
		return "author.org"
	default:
		return ""
	}
}
