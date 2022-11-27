package logic

import (
	"context"
	"soft2_backend/service/paper/api/internal/svc"
	"soft2_backend/service/paper/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaperLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaperLogic {
	return &PaperLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaperLogic) Paper(req *types.PaperRequest) (resp *types.PaperResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
