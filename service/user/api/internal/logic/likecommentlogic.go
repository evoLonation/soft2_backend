package logic

import (
	"context"
	"encoding/json"
	"soft2_backend/service/user/model"

	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikeCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeCommentLogic {
	return &LikeCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikeCommentLogic) LikeComment(req *types.LikeCommentRequest) (resp *types.LikeCommentResponse, err error) {
	// todo: add your logic here and delete this line
	userId, _ := l.ctx.Value("UserId").(json.Number).Int64()
	newLikeComment := model.Like{
		UserId:    userId,
		CommentId: req.CommentId,
	}
	comment, _ := l.svcCtx.CommentModel.FindOne(l.ctx, req.CommentId)
	comment.Likes = comment.Likes + 1
	_ = l.svcCtx.CommentModel.Update(l.ctx, comment)
	_, err = l.svcCtx.LikeModel.Insert(l.ctx, &newLikeComment)
	return &types.LikeCommentResponse{Code: 0}, nil
}
