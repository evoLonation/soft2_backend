package logic

import (
	"context"
	"encoding/json"

	"soft2_backend/service/help/api/internal/svc"
	"soft2_backend/service/help/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRequestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRequestLogic {
	return &UserRequestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRequestLogic) UserRequest(req *types.UserReqReq) (resp *types.UserReqReply, err error) {
	// todo: add your logic here and delete this line
	UserId, _ := l.ctx.Value("UserId").(json.Number).Int64()
	reqList, err := l.svcCtx.LiteratureRequestModel.FindByUser(l.ctx, UserId, req.Type)
	sum := len(reqList)
	var reql []types.UserReq
	/*for i, oneReq := range reqList {
		if i >= int(req.End) {
			break
		}
		if i >= int(req.Start) {
			var request types.UserReq
			request.RequestId = oneReq.Id
			request.RequestTime = oneReq.RequestTime.GoString()
			request.RequestContent = oneReq.RequestContent
			request.Wealth = oneReq.Wealth
			request.Type = oneReq.RequestStatus
			help, _ := l.svcCtx.LiteratureHelpModel.FindOneByReqId(l.ctx, oneReq.Id)
			request.HelpId = help.Id
			reql = append(reql, request)
		}
	}*/
	return &types.UserReqReply{
		Requests:   reql,
		RequestNum: int64(sum),
	}, nil
}
