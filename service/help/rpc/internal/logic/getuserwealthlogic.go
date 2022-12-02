package logic

import (
	"context"

	"soft2_backend/service/help/rpc/internal/svc"
	"soft2_backend/service/help/rpc/types/help"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserWealthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserWealthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserWealthLogic {
	return &GetUserWealthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserWealthLogic) GetUserWealth(in *help.WealthReq) (*help.WealthReply, error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserHelpModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &help.WealthReply{
		Wealth: user.Wealth,
	}, nil
}
