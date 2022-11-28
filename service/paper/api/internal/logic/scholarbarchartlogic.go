package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"soft2_backend/service/paper/database"

	"soft2_backend/service/paper/api/internal/svc"
	"soft2_backend/service/paper/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ScholarBarchartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScholarBarchartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScholarBarchartLogic {
	return &ScholarBarchartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScholarBarchartLogic) ScholarBarchart(req *types.ScholarBarchartRequest) (resp *types.ScholarBarchartResponse, err error) {
	// todo: add your logic here and delete this line
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"id": req.ScholarId,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Printf("Error encoding query: %s\n", err)
	}
	log.Println(buf.String())
	res := database.SearchAuthor(buf)

	source := res["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})
	resp = &types.ScholarBarchartResponse{
		Achievements: int(source["n_pubs"].(float64)),
		References:   int(source["n_citation"].(float64)),
	}
	return
}
