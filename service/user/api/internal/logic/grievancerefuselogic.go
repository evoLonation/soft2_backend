package logic

import (
	"context"
	"fmt"
	"soft2_backend/service/apply/rpc/types/apply"
	message2 "soft2_backend/service/message/rpc/types/message"
	paper2 "soft2_backend/service/paper/rpc/paper"

	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GrievanceRefuseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGrievanceRefuseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GrievanceRefuseLogic {
	return &GrievanceRefuseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GrievanceRefuseLogic) GrievanceRefuse(req *types.GrievanceRefuseRequest) (resp *types.GrievanceRefuseResponse, err error) {
	// todo: add your logic here and delete this line
	//给申诉者发通知
	grievance, err := l.svcCtx.GrievanceModel.FindOne(l.ctx, req.GrievanceId)
	plantiffId := grievance.PlaintiffId
	fmt.Printf("111111111\n%s", plantiffId)
	plaintiff, _ := l.svcCtx.ApplyRpc.CheckUser(l.ctx, &apply.CheckUserReq{ScholarId: plantiffId})
	fmt.Printf("22222222\n%d", plaintiff.UserId)
	paperId := grievance.PaperId
	paper, _ := l.svcCtx.PaperRpc.GetPaper(l.ctx, &paper2.GetPaperReq{PaperId: paperId})
	fmt.Printf("333333\n%s", paper.PaperName)
	//var papername string
	//if len(paper.PaperName) > 20 {
	//	papername = paper.PaperName[0:20] + "..."
	//} else {
	//	papername = paper.PaperName
	//}
	content := fmt.Sprintf("你对文献 %s 的申诉未通过", grievance.PaperId)
	_, err = l.svcCtx.MessageRpc.CreateMessage(l.ctx, &message2.CreateMessageReq{
		ReceiverId:  plaintiff.UserId,
		Content:     content,
		MessageType: 6,
		Result:      1,
		GId:         req.GrievanceId,
		PId:         paperId,
	})
	if err != nil {
		return nil, err
	}
	return &types.GrievanceRefuseResponse{}, nil
}
