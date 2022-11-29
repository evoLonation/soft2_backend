package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"soft2_backend/service/apply/rpc/internal/svc"
	"soft2_backend/service/apply/rpc/types/apply"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserLogic {
	return &CheckUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckUserLogic) CheckUser(in *apply.CheckUserReq) (*apply.CheckUserReply, error) {
	identify, err := l.svcCtx.ApplyModel.FindByScholarId(l.ctx, in.ScholarId)
	if err == nil {
		return &apply.CheckUserReply{
			IsVerified: true,
			UserId:     identify.UserId,
		}, nil
	} else if err == sqlc.ErrNotFound {
		return &apply.CheckUserReply{
			IsVerified: false,
		}, nil
	} else {
		return nil, err
	}
}
