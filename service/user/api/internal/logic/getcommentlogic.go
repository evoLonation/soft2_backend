package logic

import (
	"context"
	"fmt"

	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentLogic {
	return &GetCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentLogic) GetComment(req *types.GetCommentRequest) (resp *types.GetCommentReply, err error) {
	// todo: add your logic here and delete this line
	search := "'" + req.PaperId + "'"
	reqList, err := l.svcCtx.CommentModel.FindByPaperId(l.ctx, search)

	var reql []types.CommentReply
	sum := len(reqList)
	if sum == 0 {
		return &types.GetCommentReply{
			HasComment: 1,
			Comments:   nil,
		}, nil
	}
	biggest := -1
	temp := 0
	for i := 0; i < sum; i++ {
		if int(reqList[i].Likes) > biggest {
			biggest = int(reqList[i].Likes)
			temp = int(reqList[i].CommentId)
		}
	}
	firstComment, _ := l.svcCtx.CommentModel.FindOne(l.ctx, int64(temp))
	var firstCommentReply types.CommentReply
	firstCommentReply.CommentId = int64(temp)
	firstCommentReply.UserName = firstComment.UserNickname
	firstCommentReply.Content = firstComment.Content
	firstCommentReply.Date = fmt.Sprintf("%d年%d月%d日", firstComment.CreateTime.Year(), firstComment.CreateTime.Month(), firstComment.CreateTime.Day())
	firstCommentReply.Likes = firstComment.Likes
	reql = append(reql, firstCommentReply)
	for i, oneReq := range reqList {
		if i > sum {
			break
		}
		var request types.CommentReply
		if int(oneReq.CommentId) == temp {
			continue
		}
		request.CommentId = oneReq.CommentId
		request.UserName = oneReq.UserNickname
		request.UserId = oneReq.UserId
		request.Content = oneReq.Content
		request.Date = fmt.Sprintf("%d年%d月%d日", oneReq.CreateTime.Year(), oneReq.CreateTime.Month(), oneReq.CreateTime.Day())
		request.Likes = oneReq.Likes
		reql = append(reql, request)
	}
	return &types.GetCommentReply{
		HasComment: 0,
		Comments:   reql,
	}, nil
}
