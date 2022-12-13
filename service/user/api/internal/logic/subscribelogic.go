package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"soft2_backend/service/apply/rpc/types/apply"
	message2 "soft2_backend/service/message/rpc/types/message"
	"soft2_backend/service/user/model"

	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"

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
	userId, _ := l.ctx.Value("UserId").(json.Number).Int64()
	user, _ := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	scholarId := req.ScholarId
	checkUser, _ := l.svcCtx.ApplyRpc.CheckUser(l.ctx, &apply.CheckUserReq{ScholarId: scholarId})
	checkUserId := checkUser.UserId
	_, err = l.svcCtx.SubscribeModel.FindSubscribeId(l.ctx, userId, scholarId)
	if err != model.ErrNotFound {
		return &types.SubscribeResponse{Code: 1}, nil
	}
	newSubscribe := model.Subscribe{
		UserId:    userId,
		ScholarId: scholarId,
	}
	_, _ = l.svcCtx.SubscribeModel.Insert(l.ctx, &newSubscribe)
	user, _ = l.svcCtx.UserModel.FindOne(l.ctx, userId)
	user.Follows = user.Follows + 1
	_ = l.svcCtx.UserModel.Update(l.ctx, user)
	var username string
	if len(user.Nickname) > 20 {
		username = user.Nickname[0:20] + "..."
	} else {
		username = user.Nickname
	}
	content := fmt.Sprintf("%s 关注了你", username)
	_, _ = l.svcCtx.MessageRpc.CreateMessage(l.ctx, &message2.CreateMessageReq{
		ReceiverId:  checkUserId,
		Content:     content,
		MessageType: 3,
		UId:         userId,
	})
	return &types.SubscribeResponse{Code: 0}, nil
}
