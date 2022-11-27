package logic

import (
	"context"

	"help/rpc/internal/svc"
	"help/rpc/types/help"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpDateStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpDateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpDateStatusLogic {
	return &UpDateStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpDateStatusLogic) UpDateStatus(in *help.IdReq) (*help.Reply, error) {
	// todo: add your logic here and delete this line

	return &help.Reply{}, nil
}
