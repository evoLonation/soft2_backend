package logic

import (
	"context"
	"encoding/json"
	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"
	"soft2_backend/service/user/model"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectPaperLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCollectPaperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectPaperLogic {
	return &CollectPaperLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CollectPaperLogic) CollectPaper(req *types.CollectPaperRequest) (resp *types.CollectPaperResponse, err error) {
	// todo: add your logic here and delete this line

	temp, err := l.svcCtx.CollectModel.FindOneByTwo(l.ctx, 3, req.PaperId)
	if temp != nil {
		return &types.CollectPaperResponse{Code: 1}, nil
	}
	userId, _ := l.ctx.Value("UserId").(json.Number).Int64()
	newCollect := model.Collect{
		UserId:     userId,
		PaperId:    req.PaperId,
		CreateTime: time.Time{},
	}
	_, err = l.svcCtx.CollectModel.Insert(l.ctx, &newCollect)
	return &types.CollectPaperResponse{Code: 0}, nil
}