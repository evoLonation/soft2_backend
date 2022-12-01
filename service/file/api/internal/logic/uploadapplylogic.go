package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"mime/multipart"
	"os"
	"soft2_backend/service/apply/rpc/types/apply"
	"soft2_backend/service/file/api/internal/svc"
	"soft2_backend/service/file/filecommon"
	"strconv"
)

type UploadApplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	Files  []struct {
		*multipart.FileHeader
		multipart.File
	}
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
	defer func() {
		if err := os.RemoveAll(filecommon.FilePath + "temp/"); err != nil {
			panic(err)
		}
	}()
	os.MkdirAll(filecommon.FilePath+"temp/", 0x777)
	for i, file := range l.Files {
		if _, err := filecommon.CreateTempFile(file.File, file.FileHeader, "file"+strconv.Itoa(i)); err != nil {
			return err
		}
	}
	zipfilename := filecommon.NewUUid() + ".zip"
	if err := zipSource(filecommon.FilePath+"temp/", filecommon.FilePath+zipfilename); err != nil {
		return err
	}
	_, err := l.svcCtx.Apply.CreateIdentify(l.ctx, &apply.CreateIdentifyReq{UserId: filecommon.GetUserId(l.ctx), ScholarId: l.ScholarId, Url: filecommon.GetUrl(zipfilename)})
	if err != nil {
		return err
	}
	return nil
}
