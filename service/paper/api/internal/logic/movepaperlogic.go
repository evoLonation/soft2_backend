package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"soft2_backend/service/paper/database"

	"soft2_backend/service/paper/api/internal/svc"
	"soft2_backend/service/paper/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MovePaperLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMovePaperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MovePaperLogic {
	return &MovePaperLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MovePaperLogic) MovePaper(req *types.MovePaperRequest) (resp *types.MovePaperResponse, err error) {
	// todo: add your logic here and delete this line
	var paperBuf bytes.Buffer
	paperQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"_id": req.PaperId,
			},
		},
	}
	if err := json.NewEncoder(&paperBuf).Encode(paperQuery); err != nil {
		log.Printf("Error encoding query: %s\n", err)
	}
	log.Println(paperBuf.String())
	paperRes := database.SearchPaper(paperBuf)

	var ownerBuf bytes.Buffer
	ownerQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"_id": req.OwnerId,
			},
		},
	}
	if err := json.NewEncoder(&ownerBuf).Encode(ownerQuery); err != nil {
		log.Printf("Error encoding query: %s\n", err)
	}
	log.Println(ownerBuf.String())
	ownerRes := database.SearchAuthor(ownerBuf)

	var targetBuf bytes.Buffer
	targetQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"_id": req.TargetId,
			},
		},
	}
	if err := json.NewEncoder(&targetBuf).Encode(targetQuery); err != nil {
		log.Printf("Error encoding query: %s\n", err)
	}
	log.Println(targetBuf.String())
	targetRes := database.SearchAuthor(targetBuf)

	paperSource := paperRes["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})
	ownerSource := ownerRes["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})
	targetSource := targetRes["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})

	paperAuthors := paperSource["authors"].([]interface{})
	for i, author := range paperAuthors {
		author := author.(map[string]interface{})
		if author["id"] == req.OwnerId {
			paperAuthors[i] = map[string]interface{}{
				"id":   req.TargetId,
				"name": targetSource["name"],
			}
			break
		}
	}

	var updateBuf bytes.Buffer
	updatePaper := map[string]interface{}{
		"doc": map[string]interface{}{
			"authors": paperAuthors,
		},
	}
	if err := json.NewEncoder(&updateBuf).Encode(updatePaper); err != nil {
		log.Printf("Error encoding query: %s\n", err)
	}
	log.Println(updateBuf.String())
	updatePaperRes := database.UpdatePaper(updateBuf, req.PaperId)
	log.Println(updatePaperRes)
	if int(updatePaperRes["_shards"].(map[string]interface{})["successful"].(float64)) != 1 {
		return nil, errors.New("update paper error")
	}

	var rank int
	ownerPapers := ownerSource["pubs"].([]interface{})
	for i, ownerPaper := range ownerPapers {
		p := ownerPaper.(map[string]interface{})
		if p["i"] == req.PaperId {
			rank = int(p["r"].(float64))
			ownerPapers = append(ownerPapers[:i], ownerPapers[i+1:]...)
			break
		}
	}
	var updateOwnerBuf bytes.Buffer
	updateOwner := map[string]interface{}{
		"doc": map[string]interface{}{
			"pubs": ownerPapers,
		},
	}
	if err := json.NewEncoder(&updateOwnerBuf).Encode(updateOwner); err != nil {
		log.Printf("Error encoding query: %s\n", err)
	}
	log.Println(updateOwnerBuf.String())
	updateOwnerRes := database.UpdateAuthor(updateOwnerBuf, req.OwnerId)
	if int(updateOwnerRes["_shards"].(map[string]interface{})["successful"].(float64)) != 1 {
		return nil, errors.New("update owner error")
	}

	targetPapers := targetSource["pubs"].([]interface{})
	targetPapers = append(targetPapers, map[string]interface{}{
		"i": req.PaperId,
		"r": rank,
	})
	var updateTargetBuf bytes.Buffer
	updateTarget := map[string]interface{}{
		"doc": map[string]interface{}{
			"pubs": targetPapers,
		},
	}
	if err := json.NewEncoder(&updateTargetBuf).Encode(updateTarget); err != nil {
		log.Printf("Error encoding query: %s\n", err)
	}
	log.Println(updateTargetBuf.String())
	updateTargetRes := database.UpdateAuthor(updateTargetBuf, req.TargetId)
	if int(updateTargetRes["_shards"].(map[string]interface{})["successful"].(float64)) != 1 {
		return nil, errors.New("update target error")
	}

	return &types.MovePaperResponse{}, nil
}
