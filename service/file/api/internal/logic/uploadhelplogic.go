package logic

import (
	"context"

	"soft2_backend/service/file/api/internal/svc"
	"soft2_backend/service/file/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadHelpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadHelpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadHelpLogic {
	return &UploadHelpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadHelpLogic) UploadHelp(req *types.UploadHelpReq) error {
	// todo: add your logic here and delete this line

	return nil
}
