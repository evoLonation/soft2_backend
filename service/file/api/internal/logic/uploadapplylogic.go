package logic

import (
	"context"
	"soft2_backend/service/file/api/internal/svc"
	"soft2_backend/service/file/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadApplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadApplyLogic {
	return &UploadApplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadApplyLogic) UploadApply(req *types.UploadApplyReq) error {
	// todo: add your logic here and delete this line

	return nil
}
