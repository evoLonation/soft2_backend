package logic

import (
	"context"
	"soft2_backend/service/help/api/internal/svc"
	"soft2_backend/service/help/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ComplaintDealLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewComplaintDealLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ComplaintDealLogic {
	return &ComplaintDealLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ComplaintDealLogic) ComplaintDeal(req *types.ComplaintDealReq) error {
	theReq, err := l.svcCtx.LiteratureRequestModel.FindOne(l.ctx, req.RequestId)
	if err != nil {
		return err
	}
	help, err := l.svcCtx.LiteratureHelpModel.FindOneByReqId(l.ctx, req.RequestId)
	if err != nil {
		if req.Res == 0 {
			theReq.RequestStatus = 4
		} else {
			theReq.RequestStatus = 2
		}
	} else {
		if req.Res == 0 {
			theReq.RequestStatus = 4
			help.HelpStatus = 4
		} else {
			theReq.RequestStatus = 2
			help.HelpStatus = 2
		}
		err := l.svcCtx.LiteratureHelpModel.Update(l.ctx, help)
		if err != nil {
			return err
		}
	}
	e := l.svcCtx.LiteratureRequestModel.Update(l.ctx, theReq)
	if e != nil {
		return e
	}
	return nil
}
