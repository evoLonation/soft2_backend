package logic

import (
	"context"

	"soft2_backend/service/help/api/internal/svc"
	"soft2_backend/service/help/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserHelpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserHelpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserHelpLogic {
	return &UserHelpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserHelpLogic) UserHelp(req *types.UserHelpReq) (resp *types.UserHelpReply, err error) {
	// todo: add your logic here and delete this line
	reqList, err := l.svcCtx.LiteratureHelpModel.FindByUserId(l.ctx, req.UserId, req.Type)
	sum := len(reqList)
	var reql []types.UserReq
	for i, oneReq := range reqList {
		if i >= int(req.End) {
			break
		}
		if i >= int(req.Start) {
			var request types.UserReq
			request.RequestId = oneReq.RequestId
			theReq, _ := l.svcCtx.LiteratureRequestModel.FindOne(l.ctx, oneReq.RequestId)
			request.RequestTime = theReq.RequestTime.GoString()
			request.RequestContent = theReq.RequestContent
			request.Wealth = oneReq.Wealth
			request.Type = oneReq.HelpStatus
			request.HelpId = oneReq.Id
			reql = append(reql, request)
		}
	}
	return &types.UserHelpReply{
		Requests: reql,
		HelpNum:  int64(sum),
	}, nil
}