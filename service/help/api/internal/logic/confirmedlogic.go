package logic

import (
	"context"
	"soft2_backend/service/message/rpc/types/message"

	"soft2_backend/service/help/api/internal/svc"
	"soft2_backend/service/help/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfirmedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConfirmedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfirmedLogic {
	return &ConfirmedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConfirmedLogic) Confirmed(req *types.ConfirmedReq) error {
	one, err := l.svcCtx.LiteratureRequestModel.FindOne(l.ctx, req.RequestId)
	if err != nil {
		return err
	}
	one.RequestStatus = 2
	err = l.svcCtx.LiteratureRequestModel.Update(l.ctx, one)
	if err != nil {
		return err
	}
	help, err := l.svcCtx.LiteratureHelpModel.FindOneByReqId(l.ctx, req.RequestId)
	if err != nil {
		return err
	}
	help.HelpStatus = 2
	err = l.svcCtx.LiteratureHelpModel.Update(l.ctx, help)
	_, err = l.svcCtx.MessageRpc.CreateMessage(l.ctx, &message.CreateMessageReq{
		RId:         req.RequestId,
		MessageType: 8,
	})
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}
