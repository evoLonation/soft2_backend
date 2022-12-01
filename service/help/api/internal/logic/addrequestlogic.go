package logic

import (
	"context"
	"encoding/json"
	"errors"
	"soft2_backend/service/help/model"
	"time"

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
	UserId, _ := l.ctx.Value("UserId").(json.Number).Int64()
	var authors string
	var n = len(req.Author)
	for i, author := range req.Author {
		authors += author
		if i+1 < n {
			authors += " "
		}
	}

	var newRequest = new(model.LiteratureRequest)
	newRequest.UserId = UserId
	newRequest.Title = req.Title
	newRequest.Author = authors
	newRequest.Magazine = req.Magazine
	newRequest.Link = req.Link
	newRequest.RequestContent = req.Content
	newRequest.Wealth = req.Wealth
	newRequest.RequestTime = time.Now()
	_, err = l.svcCtx.LiteratureRequestModel.Insert(l.ctx, newRequest)
	if err != nil {
		return nil, err
	}
	user, err := l.svcCtx.UserHelpModel.FindOne(l.ctx, UserId)
	if err != nil {
		return nil, err
	}
	user.Request += 1
	if user.Wealth < req.Wealth {
		err := errors.New("财富之不足")
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	user.Wealth -= req.Wealth
	err = l.svcCtx.UserHelpModel.Update(l.ctx, user)
	if err != nil {
		return nil, err
	}
	return &types.AddRequestsReply{}, nil
}
