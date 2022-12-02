package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"soft2_backend/service/user/model"

	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStarPaperLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStarPaperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStarPaperLogic {
	return &GetStarPaperLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStarPaperLogic) GetStarPaper(req *types.GetStarPaperRequest) (resp *types.GetStarPaperResponse, err error) {
	// todo: add your logic here and delete this line
	userId, _ := l.ctx.Value("UserId").(json.Number).Int64()
	reqList, err := l.svcCtx.CollectModel.FindByUserId(l.ctx, userId)
	if err == model.ErrNotFound {
		return &types.GetStarPaperResponse{PaperStar: nil}, nil
	}
	var reql []types.PaperStarReply
	sum := len(reqList)
	for i, onReq := range reqList {
		if i > sum {
			break
		}
		var request types.PaperStarReply
		request.PaperId = onReq.PaperId
		request.Date = fmt.Sprintf("%d年%d月%d日", onReq.CreateTime.Year(), onReq.CreateTime.Month(), onReq.CreateTime.Day())
		reql = append(reql, request)
	}
	return &types.GetStarPaperResponse{PaperStar: reql}, nil
}
