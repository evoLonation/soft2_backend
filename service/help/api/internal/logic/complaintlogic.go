package logic

import (
	"context"

	"soft2_backend/service/help/api/internal/svc"
	"soft2_backend/service/help/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ComplaintLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewComplaintLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ComplaintLogic {
	return &ComplaintLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ComplaintLogic) Complaint(req *types.ComplaintReq) error {
	// todo: add your logic here and delete this line
	request, err := l.svcCtx.LiteratureRequestModel.FindOne(l.ctx, req.RequestId)
	if err != nil {
		return err
	}
	request.Complaint = req.Content
	request.RequestStatus = 3
	err = l.svcCtx.LiteratureRequestModel.Update(l.ctx, request)
	if err != nil {
		return err
	}
	help, err := l.svcCtx.LiteratureHelpModel.FindOneByReqId(l.ctx, req.RequestId)
	if err != nil {
		return err
	}
	help.HelpStatus = 3
	err = l.svcCtx.LiteratureHelpModel.Update(l.ctx, help)
	if err != nil {
		return err
	}
	return nil
}
