package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"soft2_backend/service/apply/rpc/applyclient"
	"soft2_backend/service/paper/database"

	"soft2_backend/service/paper/api/internal/svc"
	"soft2_backend/service/paper/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ScholarClaimLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScholarClaimLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScholarClaimLogic {
	return &ScholarClaimLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScholarClaimLogic) ScholarClaim(req *types.ScholarClaimRequest) (resp *types.ScholarClaimResponse, err error) {
	// todo: add your logic here and delete this line
	var paperBuf bytes.Buffer
	paperQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"id": req.PaperId,
			},
		},
	}
	if err := json.NewEncoder(&paperBuf).Encode(paperQuery); err != nil {
		log.Printf("Error encoding query: %s", err)
	}
	log.Println(paperBuf.String())
	paperRes := database.SearchPaper(paperBuf)

	var scholarBuf bytes.Buffer
	scholarQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"id": req.ScholarId,
			},
		},
	}
	if err := json.NewEncoder(&scholarBuf).Encode(scholarQuery); err != nil {
		log.Printf("Error encoding query: %s", err)
	}
	log.Println(scholarBuf.String())
	scholarRes := database.SearchAuthor(scholarBuf)

	paperSource := paperRes["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})
	scholarSource := scholarRes["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})

	minLevenshtein := 100.0
	minIter := 0
	authors := NilHandler(paperSource["authors"], "list").([]interface{})
	for i, author := range authors {
		if NilHandler(author.(map[string]interface{})["id"], "string").(string) == req.ScholarId {
			return &types.ScholarClaimResponse{
				Code:      0,
				ScholarId: "",
			}, nil
		}
		authorName := NilHandler(author.(map[string]interface{})["name"], "string").(string)
		thisLevenshtein := Levenshtein(authorName, scholarSource["name"].(string), 1, 1, 1)
		if thisLevenshtein < minLevenshtein {
			minLevenshtein = thisLevenshtein
			minIter = i
		}
	}

	minLevenAuthor := authors[minIter].(map[string]interface{})
	checkScholar, err := l.svcCtx.ApplyRpc.CheckUser(l.ctx, &applyclient.CheckUserReq{
		ScholarId: minLevenAuthor["id"].(string),
	})
	if err != nil {
		return nil, err
	}

	if !checkScholar.IsVerified {
		oldScholarId := authors[minIter].(map[string]interface{})["id"].(string)
		authors[minIter].(map[string]interface{})["id"] = req.ScholarId
		authors[minIter].(map[string]interface{})["name"] = scholarSource["name"]

		var updatePaperBuf bytes.Buffer
		updatePaper := map[string]interface{}{
			"doc": map[string]interface{}{
				"authors": authors,
			},
		}
		if err := json.NewEncoder(&updatePaperBuf).Encode(updatePaper); err != nil {
			log.Printf("Error encoding query: %s\n", err)
		}
		log.Println(updatePaperBuf.String())
		updatePaperRes := database.UpdatePaper(updatePaperBuf, req.PaperId)
		if int(updatePaperRes["_shards"].(map[string]interface{})["successful"].(float64)) != 1 {
			return nil, errors.New("update paper error")
		}

		var updateNewScholarBuf bytes.Buffer
		scholarPapers := scholarSource["pubs"].([]interface{})
		scholarPapers = append(scholarPapers, map[string]interface{}{
			"i": req.PaperId,
			"r": minIter,
		})
		updateScholar := map[string]interface{}{
			"doc": map[string]interface{}{
				"pubs":   scholarPapers,
				"n_pubs": len(scholarPapers),
			},
		}
		if err := json.NewEncoder(&updateNewScholarBuf).Encode(updateScholar); err != nil {
			log.Printf("Error encoding query: %s\n", err)
		}
		log.Println(updateNewScholarBuf.String())
		updateScholarRes := database.UpdateAuthor(updateNewScholarBuf, req.ScholarId)
		if int(updateScholarRes["_shards"].(map[string]interface{})["successful"].(float64)) != 1 {
			return nil, errors.New("update scholar error")
		}

		var findOldScholarBuf bytes.Buffer
		findOldScholarQuery := map[string]interface{}{
			"query": map[string]interface{}{
				"match": map[string]interface{}{
					"id": oldScholarId,
				},
			},
		}
		if err := json.NewEncoder(&findOldScholarBuf).Encode(findOldScholarQuery); err != nil {
			log.Printf("Error encoding query: %s", err)
		}
		log.Println(findOldScholarBuf.String())
		findOldScholarRes := database.SearchAuthor(findOldScholarBuf)

		oldScholarHits := NilHandler(findOldScholarRes["hits"].(map[string]interface{})["hits"], "list").([]interface{})
		if len(oldScholarHits) != 0 {
			oldScholarSource := oldScholarHits[0].(map[string]interface{})["_source"].(map[string]interface{})
			oldScholarPapers := NilHandler(oldScholarSource["pubs"], "list").([]interface{})
			for i, paper := range oldScholarPapers {
				if paper.(map[string]interface{})["i"].(string) == req.PaperId {
					oldScholarPapers = append(oldScholarPapers[:i], oldScholarPapers[i+1:]...)
					break
				}
			}
			var updateOldScholarBuf bytes.Buffer
			updateOldScholar := map[string]interface{}{
				"doc": map[string]interface{}{
					"pubs":   oldScholarPapers,
					"n_pubs": len(oldScholarPapers),
				},
			}
			if err := json.NewEncoder(&updateOldScholarBuf).Encode(updateOldScholar); err != nil {
				log.Printf("Error encoding query: %s\n", err)
			}
			log.Println(updateOldScholarBuf.String())
			updateOldScholarRes := database.UpdateAuthor(updateNewScholarBuf, req.ScholarId)
			if int(updateOldScholarRes["_shards"].(map[string]interface{})["successful"].(float64)) != 1 {
				return nil, errors.New("update scholar error")
			}
		}

		return &types.ScholarClaimResponse{
			Code:      0,
			ScholarId: "",
		}, nil
	} else {
		return &types.ScholarClaimResponse{
			Code:      1,
			ScholarId: NilHandler(minLevenAuthor["id"], "string").(string),
		}, nil
	}
}
