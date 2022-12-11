package logic

import (
	"context"
	"soft2_backend/service/apply/api/internal/svc"
	"soft2_backend/service/apply/api/internal/types"
	"soft2_backend/service/message/rpc/types/message"

	"github.com/zeromicro/go-zero/core/logx"
)

type DealApplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDealApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DealApplyLogic {
	return &DealApplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DealApplyLogic) DealApply(req *types.DealApplyRequest) error {
	apply, err := l.svcCtx.ApplyModel.FindOne(l.ctx, req.ApplyId)
	if err != nil {
		return err
	}
	if req.IsAgree {
		apply.Status = 1
	} else {
		apply.Status = 2
	}
	err = l.svcCtx.ApplyModel.Update(l.ctx, apply)
	if err != nil {
		return err
	}
	if req.IsAgree {
		_, err = l.svcCtx.MessageRpc.CreateMessage(l.ctx, &message.CreateMessageReq{
			ReceiverId: apply.UserId,
			Content:    "你发起的学者认证通过",
			Result:     0,
		})
	} else {
		_, err = l.svcCtx.MessageRpc.CreateMessage(l.ctx, &message.CreateMessageReq{
			ReceiverId: apply.UserId,
			Content:    "你发起的学者认证未通过",
			Result:     1,
		})
	}
	if err != nil {
		return err
	}
	return nil
}
