package logic

import (
	"Ingoland/backend/mall/service/user/model"
	"context"

	"Ingoland/backend/mall/service/user/api/internal/svc"
	"Ingoland/backend/mall/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubscribeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubscribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubscribeLogic {
	return &SubscribeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubscribeLogic) Subscribe(req *types.SubscribeRequest) (resp *types.SubscribeResponse, err error) {
	// todo: add your logic here and delete this line
	userId := req.UserId
	scholarId := req.ScholarId
	temp, _ := l.svcCtx.SubscribeModel.FindSubscribeId(l.ctx, userId, scholarId)
	if temp != nil {
		return &types.SubscribeResponse{Code: 1}, nil
	}
	newSubscribe := model.Subscribe{
		UserId:    userId,
		ScholarId: scholarId,
	}
	_, _ = l.svcCtx.SubscribeModel.Insert(l.ctx, &newSubscribe)
	user, _ := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	user.Follows = user.Follows + 1
	_ = l.svcCtx.UserModel.Update(l.ctx, user)
	return &types.SubscribeResponse{Code: 0}, nil
}
