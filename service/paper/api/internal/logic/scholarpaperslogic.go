package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
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
		"from": req.Start,
		"size": req.End - req.Start,
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
	pubs := res["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})["pubs"].([]interface{})
	for _, pub := range pubs {
		if req.IsFirst && pub.(map[string]interface{})["r"] != 0 {
			continue
		}
		var paperBuf bytes.Buffer
		paperQuery := map[string]interface{}{
			"query": map[string]interface{}{
				"match": map[string]interface{}{
					"id": pub.(map[string]interface{})["i"],
				},
			},
		}
		if err := json.NewEncoder(&paperBuf).Encode(paperQuery); err != nil {
			log.Printf("Error encoding query: %s\n", err)
		}
		log.Println(paperBuf.String())
		paperRes := database.SearchPaper(paperBuf)
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
				Name:  author.(map[string]interface{})["name"].(string),
				Id:    NilHandler(author.(map[string]interface{})["id"], "string").(string),
				HasId: hasId,
			})
		}
		papers = append(papers, &types.PaperResponseJSON{
			Title:     paper["title"].(string),
			Abstract:  NilHandler(paper["abstract"], "string").(string),
			Authors:   authors,
			Year:      NilHandler(paper["year"], "int").(int),
			NCitation: NilHandler(paper["n_citation"], "int").(int),
			Publisher: NilHandler(paper["venue"].(map[string]interface{})["raw"], "string").(string),
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
		PaperNum: int(res["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		Papers:   sortedPapers,
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
