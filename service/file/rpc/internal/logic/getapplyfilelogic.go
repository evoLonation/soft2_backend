package logic

import (
	"context"
	"errors"

	"soft2_backend/service/file/rpc/internal/svc"
	"soft2_backend/service/file/rpc/types/file"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApplyFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetApplyFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApplyFileLogic {
	return &GetApplyFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetApplyFileLogic) GetApplyFile(in *file.HelpIdReq) (*file.UrlReply, error) {
	return nil, errors.New("该接口已被启用！")
}
