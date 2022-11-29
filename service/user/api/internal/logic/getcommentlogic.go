package logic

import (
	"context"
	"encoding/json"
	"soft2_backend/service/user/model"

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
	//判断是否登录
	_, err = l.ctx.Value("UserId").(json.Number).Int64()
	islog := 0 //登录为1
	if err == nil {
		islog = 1
	}
	reqList, err := l.svcCtx.CommentModel.FindByPaperId(l.ctx, req.PaperId)
	var reql []types.CommentReply
	sum := len(reqList)
	biggest := 0 //最大点赞量
	var temp = 0 //最大点赞的评论id
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
	firstCommentReply.Date = firstComment.CreateTime.GoString()
	firstCommentReply.Likes = firstComment.Likes
	if islog == 0 {
		firstCommentReply.Liked = 0
	} else {
		_, err := l.svcCtx.LikeModel.FindLikeId(l.ctx, firstComment.UserId, firstComment.CommentId)
		if err == model.ErrNotFound {
			firstCommentReply.Liked = 0
		} else {
			firstCommentReply.Liked = 1
		}
	}
	reql = append(reql, firstCommentReply)
	for i, oneReq := range reqList {
		if i >= int(req.End) {
			break
		}
		if i >= int(req.Start) {
			var request types.CommentReply
			if int(oneReq.CommentId) == temp {
				continue
			}
			request.CommentId = oneReq.CommentId
			request.UserName = oneReq.UserNickname
			request.UserId = oneReq.UserId
			request.Content = oneReq.Content
			request.Date = oneReq.CreateTime.GoString()
			request.Likes = oneReq.Likes
			if islog == 0 {
				request.Liked = 0
			} else {
				_, err := l.svcCtx.LikeModel.FindLikeId(l.ctx, oneReq.UserId, oneReq.CommentId)
				if err == model.ErrNotFound {
					request.Liked = 0
				} else {
					request.Liked = 1
				}
			}
			reql = append(reql, request)
		}
	}
	return &types.GetCommentReply{Comments: reql}, nil
}
