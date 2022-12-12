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

type Coops []*types.CoopJSON

type ScholarCooperationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScholarCooperationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScholarCooperationLogic {
	return &ScholarCooperationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScholarCooperationLogic) ScholarCooperation(req *types.ScholarCooperationRequest) (resp *types.ScholarCooperationResponse, err error) {
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
	pubs := NilHandler(source["pubs"], "list").([]interface{})
	var pubIds []string
	for _, pub := range pubs {
		pubIds = append(pubIds, NilHandler(pub.(map[string]interface{})["i"], "string").(string))
	}
	var paperBuf bytes.Buffer
	mget := map[string]interface{}{
		"ids": pubIds,
	}
	if err := json.NewEncoder(&paperBuf).Encode(mget); err != nil {
		log.Printf("Error encoding query: %s\n", err)
	}
	log.Println(paperBuf.String())
	res = database.MgetPaper(paperBuf)

	papers := NilHandler(res["docs"], "list").([]interface{})
	coopList := make(map[string]types.CoopJSON)
	paperCnt := 0
	for _, paper := range papers {
		if paper.(map[string]interface{})["found"].(bool) == false {
			continue
		}
		if paperCnt > 80 {
			break
		}
		paperCnt++

		authors := NilHandler(paper.(map[string]interface{})["_source"].(map[string]interface{})["authors"], "list").([]interface{})
		for _, author := range authors {
			authorId := NilHandler(author.(map[string]interface{})["id"], "string").(string)
			if authorId == req.ScholarId {
				continue
			}
			v, ok := coopList[author.(map[string]interface{})["id"].(string)]
			if ok {
				v.Time++
				coopList[author.(map[string]interface{})["id"].(string)] = v
			} else {
				coopJSON := types.CoopJSON{
					ScholarId: author.(map[string]interface{})["id"].(string),
					Name:      author.(map[string]interface{})["name"].(string),
					Time:      1,
				}

				var coopBuf bytes.Buffer
				authorQuery := map[string]interface{}{
					"query": map[string]interface{}{
						"match": map[string]interface{}{
							"id": author.(map[string]interface{})["id"].(string),
						},
					},
				}
				if err := json.NewEncoder(&coopBuf).Encode(authorQuery); err != nil {
					log.Printf("Error encoding query: %s\n", err)
				}
				log.Println(coopBuf.String())
				res = database.SearchAuthor(coopBuf)
				source := res["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})
				var institutions []string
				for _, institution := range source["orgs"].([]interface{}) {
					institutions = append(institutions, institution.(string))
				}
				coopJSON.Institution = institutions
				coopList[author.(map[string]interface{})["id"].(string)] = coopJSON

			}
		}
	}

	var coopJSONs Coops
	for _, v := range coopList {
		coopJSONs = append(coopJSONs, &v)
	}

	sort.Sort(sort.Reverse(coopJSONs))
	var coopListResp []types.CoopJSON
	for i := 0; i < len(coopJSONs); i++ {
		if i > 4 {
			break
		}
		coopListResp = append(coopListResp, *coopJSONs[i])
	}
	resp = &types.ScholarCooperationResponse{
		CoopList: coopListResp,
	}
	return resp, nil
}

func (p Coops) Len() int {
	return len(p)
}

func (p Coops) Less(i, j int) bool {
	return p[i].Time < p[j].Time
}

func (p Coops) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}