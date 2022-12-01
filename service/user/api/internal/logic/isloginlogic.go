package logic

import (
	"context"
	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIsLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsLoginLogic {
	return &IsLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsLoginLogic) IsLogin(req *types.IsLoginRequest) (resp *types.IsLoginResponse, err error) {
	// todo: add your logic here and delete this line

	return &types.IsLoginResponse{Code: 0}, nil
}
