package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"mime/multipart"
	"soft2_backend/service/file/api/internal/svc"
	"soft2_backend/service/file/filecommon"
	"soft2_backend/service/file/model"
)

type UploadAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	File   struct {
		*multipart.FileHeader
		multipart.File
	}
}

func NewUploadAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadAvatarLogic {
	return &UploadAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadAvatarLogic) UploadAvatar() error {
	filename, err := filecommon.CreateUUidFile(l.File.File, l.File.FileHeader)
	if err != nil {
		return err
	}
	userId := l.ctx.Value("userId").(int64)
	err = l.svcCtx.UserAvatarModel.Delete(l.ctx, userId)
	err = filecommon.SqlErrorCheck(err)
	if err != nil && err != filecommon.NoRowError {
		return err
	}
	_, err = l.svcCtx.UserAvatarModel.Insert(l.ctx, &model.UserAvatar{UserId: userId, FileName: filename})
	err = filecommon.SqlErrorCheck(err)
	if err != nil {
		return err
	}
	return nil
}
