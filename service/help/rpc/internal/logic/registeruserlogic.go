package logic

import (
	"context"

	"help/rpc/internal/svc"
	"help/rpc/types/help"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterUserLogic {
	return &RegisterUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterUserLogic) RegisterUser(in *help.IdReq) (*help.Reply, error) {
	// todo: add your logic here and delete this line

	return &help.Reply{}, nil
}
