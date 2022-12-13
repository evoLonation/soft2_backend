package logic

import (
	"context"

	"soft2_backend/service/paper/api/internal/svc"
	"soft2_backend/service/paper/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DataInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDataInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DataInfoLogic {
	return &DataInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DataInfoLogic) DataInfo(req *types.DataInfoRequest) (resp *types.DataInfoResponse, err error) {
	// todo: add your logic here and delete this line
	return nil, nil
}
