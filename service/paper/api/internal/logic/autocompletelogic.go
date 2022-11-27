package logic

import (
	"context"
	"soft2_backend/service/paper/api/internal/svc"
	"soft2_backend/service/paper/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AutoCompleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAutoCompleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AutoCompleteLogic {
	return &AutoCompleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AutoCompleteLogic) AutoComplete(req *types.AutoCompleteRequest) (resp *types.AutoCompleteResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
