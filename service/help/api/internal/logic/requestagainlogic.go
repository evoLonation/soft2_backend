package logic

import (
	"context"

	"soft2_backend/service/help/api/internal/svc"
	"soft2_backend/service/help/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RequestAgainLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRequestAgainLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RequestAgainLogic {
	return &RequestAgainLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RequestAgainLogic) RequestAgain(req *types.RequestAgainReq) error {
	request, err := l.svcCtx.LiteratureRequestModel.FindOne(l.ctx, req.RequestId)
	if err != nil {
		return err
	}
	request.RequestStatus = 0
	err = l.svcCtx.LiteratureRequestModel.Update(l.ctx, request)
	if err != nil {
		return err
	}
	help, err := l.svcCtx.LiteratureHelpModel.FindOneByReqId(l.ctx, req.RequestId)
	if err != nil {
		return err
	}
	err = l.svcCtx.LiteratureHelpModel.Delete(l.ctx, help.Id)
	if err != nil {
		return err
	}
	return nil
}
