package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"soft2_backend/service/apply/rpc/types/apply"
	message2 "soft2_backend/service/message/rpc/types/message"
	"soft2_backend/service/paper/rpc/paper"
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
	userId, _ := l.ctx.Value("UserId").(json.Number).Int64()
	user, _ := l.svcCtx.UserModel.FindOne(l.ctx, userId) //发表评论的用户
	newComment := model.Comment{
		UserId:       userId,
		UserNickname: user.Nickname,
		PaperId:      req.PaperId,
		Content:      req.Content,
		Likes:        0,
		CreateTime:   time.Time{},
	}
	_, err = l.svcCtx.CommentModel.Insert(l.ctx, &newComment)
	//通知
	getPaper, _ := l.svcCtx.PaperRpc.GetPaper(l.ctx, &paper.GetPaperReq{PaperId: req.PaperId})
	sum := len(getPaper.Authors)
	var username string
	if len(user.Nickname) > 20 {
		username = user.Nickname[0:20] + "..."
	} else {
		username = user.Nickname
	}
	for i := 0; i < sum; i++ {
		if getPaper.Authors[i].HasId == false {
			continue
		}
		tempUser, _ := l.svcCtx.ApplyRpc.CheckUser(l.ctx, &apply.CheckUserReq{ScholarId: getPaper.Authors[i].Id})
		if tempUser.IsVerified == false {
			continue
		}
		content := fmt.Sprintf("%s评论了你的文献", username)
		_, _ = l.svcCtx.MessageRpc.CreateMessage(l.ctx, &message2.CreateMessageReq{
			ReceiverId:  tempUser.UserId,
			Content:     content,
			MessageType: 1,
			UId:         userId,
			PId:         req.PaperId,
		})
	}

	return &types.CommentPaperResponse{Code: 0}, nil
}
