package logic

import (
	"context"
	"soft2_backend/service/paper/api/internal/svc"
	"soft2_backend/service/paper/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaperCiteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaperCiteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaperCiteLogic {
	return &PaperCiteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaperCiteLogic) PaperCite(req *types.PaperCiteRequest) (resp *types.PaperCiteResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
