package logic

import (
	"context"
	"encoding/json"
	"soft2_backend/service/user/model"

	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelLikeLogic {
	return &CancelLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelLikeLogic) CancelLike(req *types.CancelLikeRequest) (resp *types.CancelLikeResponse, err error) {
	// todo: add your logic here and delete this line
	userId, _ := l.ctx.Value("UserId").(json.Number).Int64()
	_, err = l.svcCtx.LikeModel.FindLikeId(l.ctx, userId, req.CommentId)
	if err == model.ErrNotFound {
		return &types.CancelLikeResponse{Code: 1}, nil
	}
	comment, _ := l.svcCtx.CommentModel.FindOne(l.ctx, req.CommentId)
	comment.Likes = comment.Likes - 1
	_ = l.svcCtx.CommentModel.Update(l.ctx, comment)
	temp, _ := l.svcCtx.LikeModel.FindLikeId(l.ctx, userId, req.CommentId)
	_ = l.svcCtx.LikeModel.Delete(l.ctx, temp.LikeId)
	return &types.CancelLikeResponse{Code: 0}, nil
}
