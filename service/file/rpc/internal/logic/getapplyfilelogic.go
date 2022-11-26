package logic

import (
	"context"

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

func (l *GetApplyFileLogic) GetApplyFile(in *file.UserIdReq) (*file.UrlReply, error) {
	// todo: add your logic here and delete this line

	return &file.UrlReply{Url: "hello, world!"}, nil
}
