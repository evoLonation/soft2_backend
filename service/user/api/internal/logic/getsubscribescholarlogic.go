package logic

import (
	"context"
	"encoding/json"
	"soft2_backend/service/user/model"

	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubscribeScholarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSubscribeScholarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubscribeScholarLogic {
	return &GetSubscribeScholarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSubscribeScholarLogic) GetSubscribeScholar(req *types.GetSubscribeScholarRequest) (resp *types.GetSubscribeScholarResponse, err error) {
	// todo: add your logic here and delete this line
	userId, _ := l.ctx.Value("UserId").(json.Number).Int64()
	reqList, err := l.svcCtx.SubscribeModel.FindByUserId(l.ctx, userId)
	if err == model.ErrNotFound {
		return &types.GetSubscribeScholarResponse{ScholarSubscribe: nil}, nil
	}
	var reql []types.ScholarSubscribeReply
	sum := len(reqList)
	for i, oneReq := range reqList {
		if i > sum {
			break
		}
		var request types.ScholarSubscribeReply
		request.ScholarId = oneReq.ScholarId
		reql = append(reql, request)
	}
	return &types.GetSubscribeScholarResponse{ScholarSubscribe: reql}, nil
}
