package logic

import (
	"context"
	"encoding/json"

	"soft2_backend/service/help/api/internal/svc"
	"soft2_backend/service/help/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoReply, err error) {
	// todo: add your logic here and delete this line
	UserId, _ := l.ctx.Value("UserId").(json.Number).Int64()
	user, err := l.svcCtx.UserHelpModel.FindOne(l.ctx, UserId)
	return &types.UserInfoReply{
		Request: user.Request,
		Help:    user.Help,
		Wealth:  user.Wealth,
	}, err
}
