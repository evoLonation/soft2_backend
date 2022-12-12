package logic

import (
	"context"
	"encoding/json"
	"soft2_backend/service/user/model"

	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IfLikedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIfLikedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IfLikedLogic {
	return &IfLikedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IfLikedLogic) IfLiked(req *types.IsLikeCommentRequest) (resp *types.IsLikeCommentResponse, err error) {
	// todo: add your logic here and delete this line
	search := "'" + req.PaperId + "'"
	userId, _ := l.ctx.Value("UserId").(json.Number).Int64()
	reqList, err := l.svcCtx.CommentModel.FindByPaperId(l.ctx, search)
	var reql []types.CommentLikedReply
	sum := len(reqList)
	if sum == 0 {
		return &types.IsLikeCommentResponse{CommentLiked: nil}, nil
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
	var firstCommentReply types.CommentLikedReply
	firstCommentId := firstComment.CommentId
	_, err = l.svcCtx.LikeModel.FindLikeId(l.ctx, userId, firstCommentId)
	if err == model.ErrNotFound {
		firstCommentReply.IsLiked = 1 //	没有点赞
	} else {
		firstCommentReply.IsLiked = 0
	}
	reql = append(reql, firstCommentReply)
	for i, oneReq := range reqList {
		if i > sum {
			break
		}
		var request types.CommentLikedReply
		if int(oneReq.CommentId) == temp {
			continue
		}
		_, err = l.svcCtx.LikeModel.FindLikeId(l.ctx, userId, oneReq.CommentId)
		if err == model.ErrNotFound {
			request.IsLiked = 1
		} else {
			request.IsLiked = 0
		}
		reql = append(reql, request)
	}
	return &types.IsLikeCommentResponse{CommentLiked: reql}, nil
}
