package logic

import (
	"context"
	"encoding/json"
	"soft2_backend/service/user/model"

	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IfCollectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIfCollectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IfCollectLogic {
	return &IfCollectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IfCollectLogic) IfCollect(req *types.IfCollectPaperRequest) (resp *types.IfCollectPaperResponse, err error) {
	// todo: add your logic here and delete this line
	paperId := req.PaperId
	userId, _ := l.ctx.Value("UserId").(json.Number).Int64()
	_, err = l.svcCtx.CollectModel.FindOneByTwo(l.ctx, userId, paperId)
	if err == model.ErrNotFound {
		return &types.IfCollectPaperResponse{IsStar: 1}, nil
	} else {
		return &types.IfCollectPaperResponse{IsStar: 0}, nil
	}
}
