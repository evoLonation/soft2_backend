package logic

import (
	"context"
	"soft2_backend/service/file/filecommon"

	"soft2_backend/service/file/rpc/internal/svc"
	"soft2_backend/service/file/rpc/types/file"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAvatarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAvatarLogic {
	return &GetUserAvatarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAvatarLogic) GetUserAvatar(in *file.UserIdReq) (*file.UrlReply, error) {
	userAvatar, err := l.svcCtx.UserAvatarModel.FindOne(l.ctx, in.Id)
	err = filecommon.SqlErrorCheck(err)
	if err != nil && err != filecommon.NoRowError {
		return nil, err
	}
	if err == filecommon.NoRowError {
		return &file.UrlReply{Url: filecommon.GetDefaultAvatarUrl()}, nil
	} else {
		return &file.UrlReply{Url: filecommon.GetUrl(userAvatar.FileName)}, nil
	}

}
