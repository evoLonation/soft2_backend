package logic

import (
	"context"

	"soft2_backend/service/help/api/internal/svc"
	"soft2_backend/service/help/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BeforeHelpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBeforeHelpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BeforeHelpLogic {
	return &BeforeHelpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BeforeHelpLogic) BeforeHelp(req *types.BeforeHelpReq) (resp *types.BeforeHelpReply, err error) {
	// todo: add your logic here and delete this line
	theReq, err := l.svcCtx.LiteratureRequestModel.FindOne(l.ctx, req.RequestId)
	var ret int64
	if theReq.RequestStatus == 1 {
		ret = 0
	} else {
		ret = 1
	}
	return &types.BeforeHelpReply{
		Status: ret,
	}, nil
}
