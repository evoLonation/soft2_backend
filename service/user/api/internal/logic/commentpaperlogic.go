package logic

import (
	"context"
	"encoding/json"
	"soft2_backend/service/user/model"
	"time"

	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"

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
	userId, _ := l.ctx.Value("UserId").(json.Number).Int64()
	user, _ := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	newComment := model.Comment{
		UserId:       userId,
		UserNickname: user.Nickname,
		PaperId:      req.PaperId,
		Content:      req.Content,
		Likes:        0,
		CreateTime:   time.Time{},
	}
	_, err = l.svcCtx.CommentModel.Insert(l.ctx, &newComment)
	return &types.CommentPaperResponse{Code: 0}, nil
}
