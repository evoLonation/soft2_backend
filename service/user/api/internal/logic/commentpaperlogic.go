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
	//getPaper, _ := l.svcCtx.PaperRpc.GetPaper()
	_, err = l.svcCtx.CommentModel.Insert(l.ctx, &newComment)
	//_, _ = l.svcCtx.MessageRpc.CreateMessage(l.ctx, &message2.CreateMessageReq{
	//	ReceiverId:  0,
	//	Content:     "",
	//	MessageType: 0,
	//	Result:      0,
	//	UId:         0,
	//	GId:         0,
	//	PId:         "",
	//	RId:         0,
	//})
	return &types.CommentPaperResponse{Code: 0}, nil
}
