package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"math"
	"soft2_backend/service/paper/database"
	"sort"

	"soft2_backend/service/paper/api/internal/svc"
	"soft2_backend/service/paper/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Papers []*types.PaperResponseJSON

type ScholarPapersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScholarPapersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScholarPapersLogic {
	return &ScholarPapersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScholarPapersLogic) ScholarPapers(req *types.ScholarPapersRequest) (resp *types.ScholarPapersResponse, err error) {
	// todo: add your logic here and delete this line
	var authorBuf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"id": req.ScholarId,
			},
		},
	}
	if err := json.NewEncoder(&authorBuf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s\n", err)
	}
	log.Println(authorBuf.String())
	res := database.SearchAuthor(authorBuf)

	var papers Papers
	var maxYear = 0
	var minYear = 3000
	pubs := res["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})["pubs"].([]interface{})
	for _, pub := range pubs {
		if req.IsFirst && int(pub.(map[string]interface{})["r"].(float64)) != 0 {
			continue
		} else if !req.IsFirst && int(pub.(map[string]interface{})["r"].(float64)) == 0 {
			continue
		}
		var paperBuf bytes.Buffer
		paperQuery := map[string]interface{}{
			"query": map[string]interface{}{
				"match": map[string]interface{}{
					"id": pub.(map[string]interface{})["i"].(string),
				},
			},
		}
		if err := json.NewEncoder(&paperBuf).Encode(paperQuery); err != nil {
			log.Printf("Error encoding query: %s\n", err)
		}
		log.Println(paperBuf.String())
		paperRes := database.SearchPaper(paperBuf)
		log.Println(paperRes)
		paper := paperRes["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})
		if req.Year != 0 && NilHandler(paper["year"], "int").(int) != req.Year {
			continue
		}
		var authors []types.AuthorJSON
		for _, author := range paper["authors"].([]interface{}) {
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
		if NilHandler(paper["year"], "int").(int) > maxYear {
			maxYear = NilHandler(paper["year"], "int").(int)
		} else if NilHandler(paper["year"], "int").(int) < minYear {
			minYear = NilHandler(paper["year"], "int").(int)
		}
		papers = append(papers, &types.PaperResponseJSON{
			Title:     NilHandler(paper["title"], "string").(string),
			Abstract:  NilHandler(paper["abstract"], "string").(string),
			Authors:   authors,
			Year:      NilHandler(paper["year"], "int").(int),
			NCitation: NilHandler(paper["n_citation"], "int").(int),
			Publisher: NilHandler(paper["venue"], "string").(string),
		})
	}
	if req.TimeOrder {
		sort.Sort(sort.Reverse(papers))
	}
	var sortedPapers []types.PaperResponseJSON
	for _, paper := range papers {
		sortedPapers = append(sortedPapers, *paper)
	}

	resp = &types.ScholarPapersResponse{
		PaperNum:  len(sortedPapers),
		StartYear: minYear,
		EndYear:   maxYear,
		Papers:    sortedPapers[int(math.Min(float64(req.Start), float64(len(sortedPapers)))):int(math.Min(float64(req.End), float64(len(sortedPapers))))],
	}
	return
}

func (p Papers) Len() int {
	return len(p)
}

func (p Papers) Less(i, j int) bool {
	return p[i].Year < p[j].Year
}

func (p Papers) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
