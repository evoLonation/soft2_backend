package logic

import (
	"context"
	"soft2_backend/service/file/rpc/types/file"

	"github.com/zeromicro/go-zero/core/logx"
	"soft2_backend/service/file/rpctest/internal/svc"
)

type RpcTestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRpcTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RpcTestLogic {
	return &RpcTestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RpcTestLogic) RpcTest() error {
	applyFile, err := l.svcCtx.FileRpc.GetApplyFile(l.ctx, &file.HelpIdReq{Id: 123})
	if err != nil {
		return err
	}
	print(applyFile)
	return nil
}
