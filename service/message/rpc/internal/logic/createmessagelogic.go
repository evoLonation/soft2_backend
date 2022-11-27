package logic

import (
	"context"
	"soft2_backend/service/message/rpc/internal/svc"
	"soft2_backend/service/message/rpc/message"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMessageLogic {
	return &CreateMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateMessageLogic) CreateMessage(in *message.CreateMessageReq) (*message.CreateMessageReply, error) {
	// todo: add your logic here and delete this line

	return &message.CreateMessageReply{}, nil
}
