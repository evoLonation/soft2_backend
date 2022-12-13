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

type DataInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDataInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DataInfoLogic {
	return &DataInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DataInfoLogic) DataInfo(req *types.DataInfoRequest) (resp *types.DataInfoResponse, err error) {
	// todo: add your logic here and delete this line
	var paperBuf bytes.Buffer
	paperQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
		"size":             0,
		"track_total_hits": true,
	}
	if err := json.NewEncoder(&paperBuf).Encode(paperQuery); err != nil {
		log.Printf("Error encoding query: %s", err)
	}
	log.Println(paperBuf.String())
	paperRes := database.SearchPaper(paperBuf)
	paperNum := int(paperRes["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))

	var scholarBuf bytes.Buffer
	scholarQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
		"size":             0,
		"track_total_hits": true,
	}
	if err := json.NewEncoder(&scholarBuf).Encode(scholarQuery); err != nil {
		log.Printf("Error encoding query: %s", err)
	}
	log.Println(scholarBuf.String())
	scholarRes := database.SearchAuthor(scholarBuf)
	scholarNum := int(scholarRes["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))

	resp = &types.DataInfoResponse{
		PaperNum:   paperNum,
		ScholarNum: scholarNum,
	}
	return resp, nil
}
