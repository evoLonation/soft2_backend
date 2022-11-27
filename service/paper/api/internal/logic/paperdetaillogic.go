package logic

import (
	"context"
	"soft2_backend/service/paper/api/internal/svc"
	"soft2_backend/service/paper/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaperDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaperDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaperDetailLogic {
	return &PaperDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaperDetailLogic) PaperDetail(req *types.PaperDetailRequest) (resp *types.PaperDetailResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
