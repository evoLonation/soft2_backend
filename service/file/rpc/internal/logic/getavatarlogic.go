package logic

import (
	"context"

	"soft2_backend/service/file/rpc/internal/svc"
	"soft2_backend/service/file/rpc/types/file"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAvatarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAvatarLogic {
	return &GetAvatarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAvatarLogic) GetAvatar(in *file.UserIdReq) (*file.UrlReply, error) {
	// todo: add your logic here and delete this line

	return &file.UrlReply{}, nil
}
