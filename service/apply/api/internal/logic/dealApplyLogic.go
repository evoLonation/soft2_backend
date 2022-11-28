package logic

import (
	"context"
	"go-zero-share/apply/api/internal/svc"
	"go-zero-share/apply/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DealApplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDealApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DealApplyLogic {
	return &DealApplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DealApplyLogic) DealApply(req *types.DealApplyRequest) error {
	apply, err := l.svcCtx.ApplyModel.FindOne(l.ctx, req.ApplyId)
	if err != nil {
		return err
	}
	if req.IsAgree {
		apply.Status = 1
	} else {
		apply.Status = 2
	}
	err = l.svcCtx.ApplyModel.Update(l.ctx, apply)
	if err != nil {
		return err
	}
	return nil
}
