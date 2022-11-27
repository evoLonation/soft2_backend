package logic

import (
	"Ingoland/backend/mall/service/user/model"
	"context"
	"time"

	"Ingoland/backend/mall/service/user/api/internal/svc"
	"Ingoland/backend/mall/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentPaperLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentPaperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentPaperLogic {
	return &CommentPaperLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentPaperLogic) CommentPaper(req *types.CommentPaperRequest) (resp *types.CommentPaperResponse, err error) {
	// todo: add your logic here and delete this line
	newComment := model.Comment{
		UserId:     3,
		PaperId:    req.PaperId,
		Content:    req.Content,
		Likes:      0,
		CreateTime: time.Time{},
	}
	_, err = l.svcCtx.CommentModel.Insert(l.ctx, &newComment)
	return &types.CommentPaperResponse{Code: 0}, nil
}
