package logic

import (
	"context"
	"soft2_backend/common"
	"soft2_backend/service/file/rpc/types/file"
	"soft2_backend/service/jwttest/api/internal/svc"
	"soft2_backend/service/jwttest/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetJwtLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetJwtLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetJwtLogic {
	return &GetJwtLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetJwtLogic) GetJwt(req *types.Request) (resp *types.Response, err error) {
	token, err := common.GetJwt(req.UserId, l.svcCtx.Config.Auth.AccessSecret)
	nickname, err := l.svcCtx.FileRpc.GetApplyFile(l.ctx, &file.HelpIdReq{Id: 123})
	resp = &types.Response{Code: 0, Token: token, UserId: req.UserId, NickName: nickname.Url}
	return
}
