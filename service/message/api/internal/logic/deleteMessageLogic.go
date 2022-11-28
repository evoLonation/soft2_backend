package logic

import (
	"context"

	"go-zero-share/message/api/internal/svc"
	"go-zero-share/message/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMessageLogic {
	return &DeleteMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMessageLogic) DeleteMessage(req *types.DeleteMessageRequest) error {
	err := l.svcCtx.MessageModel.Delete(l.ctx, req.MessageId)
	if err != nil {
		return err
	}
	return nil
}
