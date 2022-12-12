package logic

import (
	"context"
	model2 "soft2_backend/service/help/model"
	"soft2_backend/service/user/model"

	"soft2_backend/service/help/rpc/internal/svc"
	"soft2_backend/service/help/rpc/types/help"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpDateStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpDateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpDateStatusLogic {
	return &UpDateStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpDateStatusLogic) UpDateStatus(in *help.UpdateReq) (*help.UpdateReply, error) {
	// todo: add your logic here and delete this line
	//求助表状态更新 应助表状态更新 用户表状态更新
	theRequest, _ := l.svcCtx.LiteratureRequestModel.FindOne(l.ctx, in.RequestId)
	theRequest.RequestStatus = in.Status
	err := l.svcCtx.LiteratureRequestModel.Update(l.ctx, theRequest)
	if err != nil {
		return nil, err
	}
	theHelp, err := l.svcCtx.LiteratureHelpModel.FindOneByReqId(l.ctx, in.RequestId)
	if err == model.ErrNotFound {
		var newHelp model2.LiteratureHelp
		newHelp.HelpStatus = in.Status
		newHelp.RequestId = in.RequestId
		newHelp.UserId = in.UserId
		newHelp.Wealth = theRequest.Wealth
		_, err2 := l.svcCtx.LiteratureHelpModel.Insert(l.ctx, &newHelp)
		if err2 != nil {
			return nil, err2
		}
	} else {
		theHelp.HelpStatus = in.Status
		err = l.svcCtx.LiteratureHelpModel.Update(l.ctx, theHelp)
		if err != nil {
			return nil, err
		}
	}
	user, err := l.svcCtx.UserHelpModel.FindOne(l.ctx, in.UserId)
	if err == model.ErrNotFound {
		var newUser = new(model2.UserHelp)
		newUser.UserId = in.UserId
		newUser.Help = 1
		newUser.Wealth = 100 + theRequest.Wealth
		newUser.Request = 0
		_, err = l.svcCtx.UserHelpModel.Insert(l.ctx, newUser)
	} else {
		user.Help += 1
		user.Wealth += theRequest.Wealth
		err = l.svcCtx.UserHelpModel.Update(l.ctx, user)
	}
	//_, err = l.svcCtx.MessageRpc.CreateMessage(l.ctx, &message.CreateMessageReq{
	//	UId:         theRequest.UserId,
	//	RId:         theRequest.Id,
	//	Content:     "你发起的文献互助收到援助",
	//	MessageType: 7,
	//})
	if err != nil {
		return nil, err
	}
	return &help.UpdateReply{
		UserId: theRequest.UserId,
	}, nil
}
