package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"soft2_backend/search/database"
	"soft2_backend/service/paper/rpc/internal/svc"
	"soft2_backend/service/paper/rpc/types/paper"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckScholarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckScholarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckScholarLogic {
	return &CheckScholarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckScholarLogic) CheckScholar(in *paper.CheckScholarReq) (*paper.CreateScholarReply, error) {
	// todo: add your logic here and delete this line
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"id": in.ScholarId,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s\n", err)
	}
	log.Println(buf.String())
	res := database.SearchPaper(buf)

	hits := res["hits"].(map[string]interface{})["hits"].([]interface{})
	if len(hits) == 0 {
		return nil, nil
	}
	resp := &paper.CreateScholarReply{
		ScholarName: hits[0].(map[string]interface{})["_source"].(map[string]interface{})["name"].(string),
		Institution: hits[0].(map[string]interface{})["_source"].(map[string]interface{})["orgs"].([]interface{})[0].(string),
	}
	return resp, nil
}
