package logic

import (
	"context"

	"go-zero-share/message/api/internal/svc"
	"go-zero-share/message/api/internal/types"

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
	message, err := l.svcCtx.MessageModel.FindOne(l.ctx, req.MessageId)
	if err != nil {
		return err
	}
	message.Read = true
	err = l.svcCtx.MessageModel.Update(l.ctx, message)
	if err != nil {
		return err
	}
	return nil
}
