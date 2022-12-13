package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"soft2_backend/service/paper/api/internal/svc"
	"soft2_backend/service/paper/api/internal/types"
	"soft2_backend/service/paper/database"
	"sort"
	"strconv"
	"strings"

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

func generateCompleteQuery(query string, years []int, themes []string) string {
	var yearsQuery string
	if len(years) != 0 {
		yearsQuery += "AND year:("
		yearStrs := make([]string, len(years))
		for i, year := range years {
			yearStrs[i] = strconv.Itoa(year)
		}
		yearsQuery += strings.Join(yearStrs, " OR ") + ")"
	}
	var themesQuery string
	if len(themes) != 0 {
		themesQuery += "AND keywords:("
		for i, theme := range themes {
			themes[i] = dealWord(theme, false)
		}
		themesQuery += strings.Join(themes, " OR ") + ")"
	}
	return query + yearsQuery + themesQuery
}

func dealWord(text string, isFuzzy bool) string {
	reservedCharacters := ".?+*|{}[]()\"\\#@&<>~"
	for _, char := range reservedCharacters {
		charStr := fmt.Sprintf("%c", char)
		text = strings.ReplaceAll(text, charStr, "\\"+charStr)
	}
	if isFuzzy {
		arr := strings.Split(text, " ")
		return strings.Join(arr, "~ AND")
	} else {
		return "\"" + text + "\""
	}
}

func (l *PaperLogic) Paper(req *types.PaperRequest) (resp *types.PaperResponse, err error) {
	//queryString := generateCompleteQuery(req.Query, req.Years, req.Themes)
	queryString := req.Query

	var buf bytes.Buffer
	query := map[string]interface{}{
		"from": req.Start,
		"size": req.End - req.Start,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"query_string": map[string]interface{}{
						"query": queryString,
					},
				},
			},
		},
	}
	if req.StartYear != 0 || req.EndYear != 0 || len(req.Years) != 0 || len(req.Themes) != 0 || len(req.Venues) != 0 || len(req.Institutions) != 0 {
		var filterList []map[string]interface{}
		if req.StartYear != 0 || req.EndYear != 0 {
			rangeMap := map[string]interface{}{}
			if req.StartYear != 0 {
				rangeMap["gte"] = req.StartYear
			}
			if req.EndYear != 0 {
				rangeMap["lte"] = req.EndYear
			}
			filterList = append(filterList, map[string]interface{}{
				"range": map[string]interface{}{
					"year": rangeMap,
				},
			})
		}
		for _, year := range req.Years {
			filterList = append(filterList, map[string]interface{}{
				"term": map[string]interface{}{
					"year": year,
				},
			})
		}
		for _, e := range req.Institutions {
			filterList = append(filterList, map[string]interface{}{
				"term": map[string]interface{}{
					"authors.org.filter": e,
				},
			})
		}
		for _, e := range req.Venues {
			filterList = append(filterList, map[string]interface{}{
				"term": map[string]interface{}{
					"venue.filter": e,
				},
			})
		}
		for _, e := range req.Themes {
			filterList = append(filterList, map[string]interface{}{
				"term": map[string]interface{}{
					"keywords.filter": e,
				},
			})
		}
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["filter"] = filterList
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
	if req.NeedFilterStatistics {
		query["aggs"] = map[string]interface{}{
			"years": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "year",
				},
			},
			//todo
			//"keywords": map[string]interface{}{
			//	"terms": map[string]interface{}{
			//		"field": "keywords.filter",
			//	},
			//},
			//"venues": map[string]interface{}{
			//	"terms": map[string]interface{}{
			//		"field": "venues.filter",
			//	},
			//},
			//"institutions": map[string]interface{}{
			//	"terms": map[string]interface{}{
			//		"field": "authors.org.filter",
			//	},
			//},
		}
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s\n", err)
	}
	log.Println(buf.String())
	res, err := database.SearchPaperE(buf)
	if err != nil {
		return nil, err
	}
	papers := make([]types.PaperResponseJSON, 0)

	for _, hit := range res["hits"].(map[string]interface{})["hits"].([]interface{}) {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		authors := make([]types.AuthorJSON, 0)
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
	}
	themes := make([]types.Statistic, 0)
	years := make(Sorter, 0)
	venues := make([]types.Statistic, 0)
	institutions := make([]types.Statistic, 0)
	if req.NeedFilterStatistics {
		agg := res["aggregations"].(map[string]interface{})
		for _, year := range agg["years"].(map[string]interface{})["buckets"].([]interface{}) {
			years = append(years, types.StatisticNumber{
				Name:  int(year.(map[string]interface{})["key"].(float64)),
				Count: int(year.(map[string]interface{})["doc_count"].(float64)),
			})
		}
		sort.Sort(years)
		//todo
		//for _, keyword := range agg["keywords"].(map[string]interface{})["buckets"].([]interface{}) {
		//	themes = append(themes, types.Statistic{
		//		Name:  keyword.(map[string]interface{})["key"].(string),
		//		Count: int(keyword.(map[string]interface{})["doc_count"].(float64)),
		//	})
		//}
		//for _, venue := range agg["venues"].(map[string]interface{})["buckets"].([]interface{}) {
		//	venues = append(venues, types.Statistic{
		//		Name:  venue.(map[string]interface{})["key"].(string),
		//		Count: int(venue.(map[string]interface{})["doc_count"].(float64)),
		//	})
		//}
		//for _, inst := range agg["institutions"].(map[string]interface{})["buckets"].([]interface{}) {
		//	institutions = append(institutions, types.Statistic{
		//		Name:  inst.(map[string]interface{})["key"].(string),
		//		Count: int(inst.(map[string]interface{})["doc_count"].(float64)),
		//	})
		//}
	}
	paperNum := int(res["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))
	if paperNum > 10000 {
		paperNum = 10000
	}
	resp = &types.PaperResponse{
		PaperNum:     paperNum,
		Papers:       papers,
		Themes:       themes,
		Years:        years,
		Institutions: institutions,
		Venues:       venues,
	}
	return resp, nil
}

type Sorter []types.StatisticNumber

func (p Sorter) Len() int {
	return len(p)
}
func (p Sorter) Less(i, j int) bool {
	return p[i].Name > p[j].Name
}
func (p Sorter) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
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
