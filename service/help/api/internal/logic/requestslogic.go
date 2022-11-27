package logic

import (
	"context"

	"soft2_backend/service/help/api/internal/svc"
	"soft2_backend/service/help/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RequestsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRequestsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RequestsLogic {
	return &RequestsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RequestsLogic) Requests(req *types.ReqsReq) (resp *types.ReqsReply, err error) {
	// todo: add your logic here and delete this line

	return
}
