package logic

import (
	"context"
	"encoding/json"
	"soft2_backend/service/user/model"

	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"

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

func (l *DeleteSubscribeLogic) DeleteSubscribe(req *types.DeleteSubscribeRequest) (resp *types.DeleteSubscribeResponse, err error) {
	// todo: add your logic here and delete this line
	userId, _ := l.ctx.Value("UserId").(json.Number).Int64()
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
