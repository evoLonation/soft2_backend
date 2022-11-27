package logic

import (
	"context"
	"soft2_backend/service/paper/api/internal/svc"
	"soft2_backend/service/paper/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ScholarBasicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScholarBasicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScholarBasicLogic {
	return &ScholarBasicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScholarBasicLogic) ScholarBasic(req *types.ScholarBasicRequest) (resp *types.ScholarBasicResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
