// Code generated by goctl. DO NOT EDIT!
// Source: message.proto

package server

import (
	"context"
	"soft2_backend/service/message/rpc/internal/logic"
	"soft2_backend/service/message/rpc/internal/svc"
	"soft2_backend/service/message/rpc/message"
)

type ApplyServer struct {
	svcCtx *svc.ServiceContext
	message.UnimplementedApplyServer
}

func NewApplyServer(svcCtx *svc.ServiceContext) *ApplyServer {
	return &ApplyServer{
		svcCtx: svcCtx,
	}
}

func (s *ApplyServer) CreateMessage(ctx context.Context, in *message.CreateMessageReq) (*message.CreateMessageReply, error) {
	l := logic.NewCreateMessageLogic(ctx, s.svcCtx)
	return l.CreateMessage(in)
}
