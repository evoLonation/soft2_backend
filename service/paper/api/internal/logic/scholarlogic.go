package logic

import (
	"context"
	"soft2_backend/service/paper/api/internal/svc"
	"soft2_backend/service/paper/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ScholarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScholarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScholarLogic {
	return &ScholarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScholarLogic) Scholar(req *types.ScholarRequest) (resp *types.ScholarResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
