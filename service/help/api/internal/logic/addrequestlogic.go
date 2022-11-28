package logic

import (
	"context"
	"soft2_backend/service/help/model"

	"soft2_backend/service/help/api/internal/svc"
	"soft2_backend/service/help/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRequestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRequestLogic {
	return &AddRequestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddRequestLogic) AddRequest(req *types.AddRequestsReq) (resp *types.AddRequestsReply, err error) {
	// todo: add your logic here and delete this line
	var authors string
	var n = len(req.Author)
	for i, author := range req.Author {
		authors += author
		if i+1 < n {
			authors += " "
		}
	}

	var newRequest = new(model.LiteratureRequest)
	newRequest.UserId = req.UserId
	newRequest.Title = req.Title
	newRequest.Author = authors
	newRequest.Magazine = req.Magazine
	newRequest.Link = req.Link
	newRequest.RequestContent = req.Content
	newRequest.Wealth = req.Wealth
	_, err = l.svcCtx.LiteratureRequestModel.Insert(l.ctx, newRequest)
	if err != nil {
		return nil, err
	}
	return &types.AddRequestsReply{}, nil
}
