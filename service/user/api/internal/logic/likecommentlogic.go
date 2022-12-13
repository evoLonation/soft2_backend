package logic

import (
	"context"
	"encoding/json"
	"fmt"
	message2 "soft2_backend/service/message/rpc/types/message"
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
	userId, _ := l.ctx.Value("UserId").(json.Number).Int64()
	_, err = l.svcCtx.LikeModel.FindLikeId(l.ctx, userId, req.CommentId)
	if err != model.ErrNotFound { //已经点过赞了
		return &types.LikeCommentResponse{Code: 1}, nil
	}
	comment, _ := l.svcCtx.CommentModel.FindOne(l.ctx, req.CommentId)
	comment.Likes = comment.Likes + 1
	_ = l.svcCtx.CommentModel.Update(l.ctx, comment)
	newLikeComment := model.Like{
		UserId:    userId,
		CommentId: req.CommentId,
	}
	_, _ = l.svcCtx.LikeModel.Insert(l.ctx, &newLikeComment)
	//发通知
	user, _ := l.svcCtx.UserModel.FindOne(l.ctx, userId) //点赞者
	paperId := comment.PaperId
	//paper, _ := l.svcCtx.PaperRpc.GetPaper(l.ctx, &paper2.GetPaperReq{PaperId: paperId})
	var username string

	if len(user.Nickname) > 20 {
		username = user.Nickname[0:20] + "..."
	} else {
		username = user.Nickname
	}
	content := fmt.Sprintf("%s 赞了你在 %s 的评论", username, paperId)
	_, _ = l.svcCtx.MessageRpc.CreateMessage(l.ctx, &message2.CreateMessageReq{
		ReceiverId:  comment.UserId,
		Content:     content,
		MessageType: 2,
		UId:         userId,
		PId:         paperId,
	})
	return &types.LikeCommentResponse{Code: 0}, nil
}
