package logic

import (
	"context"
	"encoding/json"
	"soft2_backend/service/user/model"
	"time"

	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"

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
	userId, _ := l.ctx.Value("UserId").(json.Number).Int64()
	_, err = l.svcCtx.CollectModel.FindOneByTwo(l.ctx, userId, req.PaperId)
	if err != model.ErrNotFound {
		return &types.CollectPaperResponse{Code: 1}, nil
	}
	newCollect := model.Collect{
		UserId:     userId,
		PaperId:    req.PaperId,
		CreateTime: time.Time{},
	}
	_, err = l.svcCtx.CollectModel.Insert(l.ctx, &newCollect)
	return &types.CollectPaperResponse{Code: 0}, nil
}
