package logic

import (
	"context"

	"help/api/internal/svc"
	"help/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRequestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRequestLogic {
	return &AddRequestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddRequestLogic) AddRequest(req *types.AddRequestsReq) (resp *types.AddRequestsReply, err error) {
	// todo: add your logic here and delete this line

	return
}
