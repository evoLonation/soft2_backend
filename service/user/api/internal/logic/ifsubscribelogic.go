package logic

import (
	"context"
	"encoding/json"
	"soft2_backend/service/user/model"

	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IfSubscribeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIfSubscribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IfSubscribeLogic {
	return &IfSubscribeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IfSubscribeLogic) IfSubscribe(req *types.IfSubscribeRequest) (resp *types.IfSubscribeResponse, err error) {
	// todo: add your logic here and delete this line
	userId, _ := l.ctx.Value("UserId").(json.Number).Int64()
	scholarId := req.ScholarId
	_, err = l.svcCtx.SubscribeModel.FindSubscribeId(l.ctx, userId, scholarId)
	if err == model.ErrNotFound {
		return &types.IfSubscribeResponse{Code: 0}, nil
	}
	return &types.IfSubscribeResponse{Code: 1}, nil
}
