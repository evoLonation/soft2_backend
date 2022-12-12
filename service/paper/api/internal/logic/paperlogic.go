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
	queryString := generateCompleteQuery(req.Query, req.Years, req.Themes)

	var buf bytes.Buffer
	query := map[string]interface{}{
		"from": req.Start,
		"size": req.End - req.Start,
		"query": map[string]interface{}{
			"query_string": map[string]interface{}{
				"query": queryString,
			},
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
