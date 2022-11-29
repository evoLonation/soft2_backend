package logic

import (
	"context"
	"mime/multipart"
	"soft2_backend/service/file/filecommon"
	"soft2_backend/service/file/model"
	"soft2_backend/service/help/rpc/types/help"

	"github.com/zeromicro/go-zero/core/logx"
	"soft2_backend/service/file/api/internal/svc"
)

type UploadHelpLogic struct {
	logx.Logger
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	File      *multipart.FileHeader
	RequestId int64
}

func NewUploadHelpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadHelpLogic {
	return &UploadHelpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadHelpLogic) UploadHelp() error {
	filename, err := filecommon.CreateFile(l.File)
	if err != nil {
		return err
	}
	userId := l.ctx.Value("userId").(int64)
	err = l.svcCtx.HelpFileModel.Delete(l.ctx, l.RequestId)
	err = filecommon.SqlErrorCheck(err)
	if err != nil && err != filecommon.NoRowError {
		return err
	}
	_, err = l.svcCtx.HelpFileModel.Insert(l.ctx, &model.HelpFile{HelpId: l.RequestId, FileName: filename})
	err = filecommon.SqlErrorCheck(err)
	if err != nil {
		return err
	}
	_, err = l.svcCtx.Help.UpDateStatus(l.ctx, &help.UpdateReq{Status: 1, UserId: userId, RequestId: l.RequestId})
	if err != nil {
		return err
	}
	return nil
}
