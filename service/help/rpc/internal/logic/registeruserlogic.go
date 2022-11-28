package logic

import (
	"context"
	"soft2_backend/service/help/model"

	"soft2_backend/service/help/rpc/internal/svc"
	"soft2_backend/service/help/rpc/types/help"

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
	var newUser = new(model.UserHelp)
	newUser.UserId = in.Id
	newUser.Help = 0
	newUser.Wealth = 0
	newUser.Request = 0
	_, err := l.svcCtx.UserHelpModel.Insert(l.ctx, newUser)
	if err != nil {
		return nil, err
	}
	return &help.Reply{}, nil
}
