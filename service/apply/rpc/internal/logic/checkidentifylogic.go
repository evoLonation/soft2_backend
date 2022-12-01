package logic

import (
	"context"
	"soft2_backend/service/apply/rpc/internal/svc"
	"soft2_backend/service/apply/rpc/types/apply"

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
	} else {
		return &apply.CheckIdentifyReply{
			IsScholar: false,
		}, nil
	}
}
