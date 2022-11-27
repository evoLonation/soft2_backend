package logic

import (
	"context"
	"soft2_backend/service/message/api/internal/svc"
	"soft2_backend/service/message/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReadMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReadMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReadMessageLogic {
	return &ReadMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReadMessageLogic) ReadMessage(req *types.ReadMessageRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
