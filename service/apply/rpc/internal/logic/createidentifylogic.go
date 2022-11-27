package logic

import (
	"context"
	"soft2_backend/service/apply/model"

	"soft2_backend/service/apply/rpc/internal/svc"
	"soft2_backend/service/apply/rpc/types/apply"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateIdentifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateIdentifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateIdentifyLogic {
	return &CreateIdentifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateIdentifyLogic) CreateIdentify(in *apply.CreateIdentifyReq) (*apply.CreateIdentifyReply, error) {

	newApply := model.Apply{
		UserId:    in.UserId,
		ScholarId: in.ScholarId,
		ApplyType: 2,
		Url:       in.Url,
	}

	_, err := l.svcCtx.ApplyModel.Insert(l.ctx, &newApply)
	if err != nil {
		return nil, err
	}

	return &apply.CreateIdentifyReply{}, nil
}
