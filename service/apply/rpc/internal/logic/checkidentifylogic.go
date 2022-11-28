package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"go-zero-share/apply/rpc/internal/svc"
	"go-zero-share/apply/rpc/types/apply"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckIdentifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckIdentifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckIdentifyLogic {
	return &CheckIdentifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckIdentifyLogic) CheckIdentify(in *apply.CheckIdentifyReq) (*apply.CheckIdentifyReply, error) {
	identify, err := l.svcCtx.ApplyModel.FindByUserId(l.ctx, in.UserId)
	if err == nil {
		return &apply.CheckIdentifyReply{
			IsScholar: true,
			ScholarId: identify.ScholarId,
		}, nil
	} else if err == sqlc.ErrNotFound {
		return &apply.CheckIdentifyReply{
			IsScholar: false,
		}, nil
	} else {
		return nil, err
	}
}