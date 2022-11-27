package logic

import (
	"context"
	"soft2_backend/service/message/api/internal/svc"
	"soft2_backend/service/message/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageLogic {
	return &GetMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMessageLogic) GetMessage() (resp *types.GetMessageResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
