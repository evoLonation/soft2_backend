package logic

import (
	"context"
	"errors"
	"soft2_backend/service/file/filecommon"

	"soft2_backend/service/file/rpc/internal/svc"
	"soft2_backend/service/file/rpc/types/file"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHelpFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetHelpFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHelpFileLogic {
	return &GetHelpFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetHelpFileLogic) GetHelpFile(in *file.ApplyIdReq) (*file.UrlReply, error) {
	helpFile, err := l.svcCtx.HelpFileModel.FindOne(l.ctx, in.Id)
	err = filecommon.SqlErrorCheck(err)
	if err != nil && err != filecommon.NoRowError {
		return nil, err
	}
	if err == filecommon.NoRowError {
		return nil, errors.New("指定的id没有找到文件！")
	} else {
		return &file.UrlReply{Url: filecommon.GetUrl(helpFile.FileName)}, nil
	}
}
