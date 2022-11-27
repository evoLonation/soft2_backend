package logic

import (
	"context"
	"soft2_backend/service/apply/api/internal/svc"
	"soft2_backend/service/apply/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApplyLogic {
	return &GetApplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApplyLogic) GetApply() (resp *types.GetApplyResponse, err error) {
	list, err := l.svcCtx.ApplyModel.FindAll(l.ctx)

	infoList := make([]types.ApplyInfo, 0)

	// 调用学者rpc查找scholarname和institution

	for _, item := range list {
		infoList = append(infoList, types.ApplyInfo{
			ApplyId:     item.ApplyId,
			ScholarName: "lll",
			Institution: "www",
			ApplyType:   item.ApplyType,
			Email:       item.Email,
			URL:         item.Url,
		})
	}

	return &types.GetApplyResponse{
		ApplyList: infoList,
	}, nil
}
