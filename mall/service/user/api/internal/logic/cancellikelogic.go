package logic

import (
	"context"
	"encoding/json"

	"Ingoland/backend/mall/service/user/api/internal/svc"
	"Ingoland/backend/mall/service/user/api/internal/types"

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
	commentId := req.CommentId
	userId, _ := l.ctx.Value("UserId").(json.Number).Int64()
	temp, err := l.svcCtx.LikeModel.FindLikeId(l.ctx, userId, commentId)
	_ = l.svcCtx.LikeModel.Delete(l.ctx, temp.LikeId)
	comment, _ := l.svcCtx.CommentModel.FindOne(l.ctx, commentId)
	comment.Likes = comment.Likes - 1
	_ = l.svcCtx.CommentModel.Update(l.ctx, comment)
	return &types.CancelLikeResponse{Code: 0}, nil
}
