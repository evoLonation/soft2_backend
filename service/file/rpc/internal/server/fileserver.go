// Code generated by goctl. DO NOT EDIT!
// Source: file.proto

package server

import (
	"context"

	"soft2_backend/service/file/rpc/internal/logic"
	"soft2_backend/service/file/rpc/internal/svc"
	"soft2_backend/service/file/rpc/types/file"
)

type FileServer struct {
	svcCtx *svc.ServiceContext
	file.UnimplementedFileServer
}

func NewFileServer(svcCtx *svc.ServiceContext) *FileServer {
	return &FileServer{
		svcCtx: svcCtx,
	}
}

func (s *FileServer) GetAvatar(ctx context.Context, in *file.UserIdReq) (*file.UrlReply, error) {
	l := logic.NewGetAvatarLogic(ctx, s.svcCtx)
	return l.GetAvatar(in)
}

func (s *FileServer) GetHelpFile(ctx context.Context, in *file.UserIdReq) (*file.UrlReply, error) {
	l := logic.NewGetHelpFileLogic(ctx, s.svcCtx)
	return l.GetHelpFile(in)
}

func (s *FileServer) GetApplyFile(ctx context.Context, in *file.UserIdReq) (*file.UrlReply, error) {
	l := logic.NewGetApplyFileLogic(ctx, s.svcCtx)
	return l.GetApplyFile(in)
}
