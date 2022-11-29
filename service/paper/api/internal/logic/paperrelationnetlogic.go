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

type PaperRelationNetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaperRelationNetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaperRelationNetLogic {
	return &PaperRelationNetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaperRelationNetLogic) PaperRelationNet(req *types.PaperRelationNetRequest) (resp *types.PaperRelationNetResponse, err error) {
	// todo: add your logic here and delete this line
	var thisPaperBuf bytes.Buffer
	thisPaperQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"id": req.Id,
			},
		},
	}
	if err := json.NewEncoder(&thisPaperBuf).Encode(thisPaperQuery); err != nil {
		log.Printf("encode query error: %v", err)
	}
	thisPaperRes := database.SearchPaper(thisPaperBuf)

	thisPaperSource := thisPaperRes["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})
	references := thisPaperSource["references"].([]interface{})
	referenceIds := make([]string, 0)
	for _, reference := range references {
		referenceIds = append(referenceIds, reference.(string))
	}

	return
}
