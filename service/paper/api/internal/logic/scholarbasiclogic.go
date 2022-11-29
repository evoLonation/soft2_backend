package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"soft2_backend/service/paper/api/internal/svc"
	"soft2_backend/service/paper/api/internal/types"
	"soft2_backend/service/paper/database"
	"sort"

	"github.com/zeromicro/go-zero/core/logx"
)

type ScholarBasicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScholarBasicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScholarBasicLogic {
	return &ScholarBasicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScholarBasicLogic) ScholarBasic(req *types.ScholarBasicRequest) (resp *types.ScholarBasicResponse, err error) {
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

	source := res["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})
	var tags []types.TagJSON
	for _, tag := range source["tags"].([]interface{}) {
		tags = append(tags, types.TagJSON{
			T: tag.(map[string]interface{})["t"].(string),
			W: int(tag.(map[string]interface{})["w"].(float64)),
		})
	}
	var institutions []string
	for _, institution := range source["orgs"].([]interface{}) {
		institutions = append(institutions, institution.(string))
	}
	var years []int
	for _, pub := range source["pubs"].([]interface{}) {
		var paperBuf bytes.Buffer
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"match": map[string]interface{}{
					"id": pub.(map[string]interface{})["i"].(string),
				},
			},
		}
		if err := json.NewEncoder(&paperBuf).Encode(query); err != nil {
			log.Printf("Error encoding query: %s\n", err)
		}
		log.Println(paperBuf.String())
		res := database.SearchPaper(paperBuf)
		source := res["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})
		if NilHandler(source["year"], "int") != nil {
			years = append(years, NilHandler(source["year"], "int").(int))
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(years)))

	resp = &types.ScholarBasicResponse{
		ScholarId:   source["id"].(string),
		Name:        source["name"].(string),
		Institution: institutions,
		Position:    NilHandler(source["position"], "string").(string),
		RefNum:      int(source["n_citation"].(float64)),
		AchNum:      int(source["n_pubs"].(float64)),
		HIndex:      int(source["h_index"].(float64)),
		Years:       years,
		Tags:        tags,
	}
	return
}
