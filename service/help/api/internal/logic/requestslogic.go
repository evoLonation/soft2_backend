package logic

import (
	"context"
	"errors"

	"soft2_backend/service/help/api/internal/svc"
	"soft2_backend/service/help/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RequestsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRequestsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RequestsLogic {
	return &RequestsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RequestsLogic) Requests(req *types.ReqsReq) (resp *types.ReqsReply, err error) {
	// todo: add your logic here and delete this line
	if req.Order != 0 && req.Order != 1 {
		return nil, errors.New("参数错误")
	}
	reqList, err := l.svcCtx.LiteratureRequestModel.FindAll(l.ctx, req.Order)
	sum := len(reqList)
	var reql []types.Request
	for i, oneReq := range reqList {
		if i >= int(req.End) {
			break
		}
		if i >= int(req.Start) {
			var request types.Request
			request.Id = oneReq.Id
			request.Title = oneReq.Title
			request.Author = oneReq.Author
			request.Magazine = oneReq.Magazine
			request.Link = oneReq.Link
			request.RequestTime = oneReq.RequestTime.Format("yyyy-mm-dd hh:mm")
			request.RequestContent = oneReq.RequestContent
			request.Wealth = oneReq.Wealth
			reql = append(reql, request)
		}
	}
	return &types.ReqsReply{
		ReqList:     reql,
		RequestsNum: int64(sum),
	}, nil
}
