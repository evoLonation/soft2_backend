package logic

import (
	"context"

	"soft2_backend/service/jwttest/api/internal/svc"
	"soft2_backend/service/jwttest/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckJwtLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckJwtLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckJwtLogic {
	return &CheckJwtLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckJwtLogic) CheckJwt() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = &types.Response{Token: l.ctx.Value("userId").(string)}
	return
}
