package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"go-zero-share/apply/api/internal/svc"
	"go-zero-share/apply/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserLogic {
	return &CheckUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckUserLogic) CheckUser(req *types.CheckUserRequest) (resp *types.CheckUserResponse, err error) {
	apply, err := l.svcCtx.ApplyModel.FindByScholarId(l.ctx, req.ScholarId)
	if err == nil {
		return &types.CheckUserResponse{
			Code:   0,
			UserId: apply.UserId,
		}, nil
	} else if err == sqlc.ErrNotFound {
		return &types.CheckUserResponse{
			Code: 1,
			Msg:  "该学者未被认证",
		}, nil
	} else {
		return nil, err
	}
}
