package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"go-zero-share/apply/api/internal/svc"
	"go-zero-share/apply/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckScholarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckScholarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckScholarLogic {
	return &CheckScholarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckScholarLogic) CheckScholar() (resp *types.CheckScholarResponse, err error) {
	userId, _ := l.ctx.Value("user_id").(json.Number).Int64()
	apply, err := l.svcCtx.ApplyModel.FindByUserId(l.ctx, userId)
	if err == nil {
		return &types.CheckScholarResponse{
			Code:      0,
			ScholarId: apply.ScholarId,
		}, nil
	} else if err == sqlc.ErrNotFound {
		return &types.CheckScholarResponse{
			Code: 1,
			Msg:  "该用户未认证学者",
		}, nil
	} else {
		return nil, err
	}
}
