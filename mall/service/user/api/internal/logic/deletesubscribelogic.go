package logic

import (
	"Ingoland/backend/mall/service/user/model"
	"context"

	"Ingoland/backend/mall/service/user/api/internal/svc"
	"Ingoland/backend/mall/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSubscribeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSubscribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSubscribeLogic {
	return &DeleteSubscribeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSubscribeLogic) DeleteSubscribe(req *types.DeleteSubscribeRequest) (resp *types.SubscribeResponse, err error) {
	// todo: add your logic here and delete this line
	userId := req.UserId
	scholarId := req.ScholarId
	temp, err := l.svcCtx.SubscribeModel.FindSubscribeId(l.ctx, userId, scholarId)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.DeleteSubscribeResponse{Code: 1}, nil
		}
	}
	_ = l.svcCtx.SubscribeModel.Delete(l.ctx, temp.SubscribeId)
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	user.Follows = user.Follows - 1
	_ = l.svcCtx.UserModel.Update(l.ctx, user)
	return &types.DeleteSubscribeResponse{Code: 0}, nil
}
