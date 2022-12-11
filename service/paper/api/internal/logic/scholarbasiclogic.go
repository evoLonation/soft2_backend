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
	var tags = make([]types.TagJSON, 0)
	for _, tag := range source["tags"].([]interface{}) {
		tags = append(tags, types.TagJSON{
			T: NilHandler(tag.(map[string]interface{})["t"], "string").(string),
			W: NilHandler(tag.(map[string]interface{})["w"], "int").(int),
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
		ScholarId:   NilHandler(source["id"], "string").(string),
		Name:        NilHandler(source["name"], "string").(string),
		Institution: institutions,
		Position:    NilHandler(source["position"], "string").(string),
		RefNum:      NilHandler(source["n_citation"], "int").(int),
		AchNum:      NilHandler(source["n_pubs"], "int").(int),
		HIndex:      source["h_index"].(int),
		Years:       years,
		Tags:        tags,
	}
	return
}
