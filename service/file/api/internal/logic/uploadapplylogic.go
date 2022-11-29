package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"mime/multipart"
	"soft2_backend/service/apply/rpc/types/apply"
	"soft2_backend/service/file/api/internal/svc"
	"soft2_backend/service/file/filecommon"
)

type UploadApplyLogic struct {
	logx.Logger
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	File      *multipart.FileHeader
	ScholarId int64
}

func NewUploadApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadApplyLogic {
	return &UploadApplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 学者认证的上传与其他上传不一样的是url还是存在了学者认证那里
func (l *UploadApplyLogic) UploadApply() error {
	filename, err := filecommon.CreateFile(l.File)
	if err != nil {
		return err
	}
	userId := l.ctx.Value("userId").(int64)
	_, err = l.svcCtx.Apply.CreateIdentify(l.ctx, &apply.CreateIdentifyReq{UserId: userId, ScholarId: l.ScholarId, Url: filecommon.GetUrl(filename)})
	if err != nil {
		return err
	}
	return nil
}
