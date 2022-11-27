// Code generated by goctl. DO NOT EDIT!
// Source: apply.proto

package server

import (
	"context"
	"soft2_backend/service/apply/rpc/internal/logic"
	"soft2_backend/service/apply/rpc/internal/svc"
	"soft2_backend/service/apply/rpc/types/apply"
)

type ApplyServer struct {
	svcCtx *svc.ServiceContext
	apply.UnimplementedApplyServer
}

func NewApplyServer(svcCtx *svc.ServiceContext) *ApplyServer {
	return &ApplyServer{
		svcCtx: svcCtx,
	}
}

func (s *ApplyServer) CreateIdentify(ctx context.Context, in *apply.CreateIdentifyReq) (*apply.CreateIdentifyReply, error) {
	l := logic.NewCreateIdentifyLogic(ctx, s.svcCtx)
	return l.CreateIdentify(in)
}