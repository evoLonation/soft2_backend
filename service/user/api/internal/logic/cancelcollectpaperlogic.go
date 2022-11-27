package logic

import (
	"context"

	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelCollectPaperLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelCollectPaperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelCollectPaperLogic {
	return &CancelCollectPaperLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelCollectPaperLogic) CancelCollectPaper(req *types.CancelCollectPaperRequest) (resp *types.CancelCollectPaperResponse, err error) {
	// todo: add your logic here and delete this line
	temp, err := l.svcCtx.CollectModel.FindOneByTwo(l.ctx, 3, req.PaperId)
	_ = l.svcCtx.CollectModel.Delete(l.ctx, temp.CollectId)
	return &types.CancelCollectPaperResponse{Code: 0}, nil
}
