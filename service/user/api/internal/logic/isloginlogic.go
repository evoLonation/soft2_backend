package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"soft2_backend/service/user/api/internal/svc"
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

func (l *IsLoginLogic) IsLogin() error {
	// todo: add your logic here and delete this line

	return nil
}
