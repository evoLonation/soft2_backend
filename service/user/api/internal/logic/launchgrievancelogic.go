package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"soft2_backend/service/apply/rpc/types/apply"
	message2 "soft2_backend/service/message/rpc/types/message"
	paper2 "soft2_backend/service/paper/rpc/paper"
	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"
	"soft2_backend/service/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type LaunchGrievanceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLaunchGrievanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LaunchGrievanceLogic {
	return &LaunchGrievanceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LaunchGrievanceLogic) LaunchGrievance(req *types.LaunchGrievanceRequest) (resp *types.LaunchGrievanceResponse, err error) {
	// todo: add your logic here and delete this line
	paperId := req.PaperId       //文献id
	defendantId := req.ScholarId //被申诉学者id
	tempId, _ := l.ctx.Value("UserId").(json.Number).Int64()
	defendUser, _ := l.svcCtx.UserModel.FindOne(l.ctx, tempId)
	username := defendUser.LoginId
	plaintiff, _ := l.svcCtx.ApplyRpc.CheckIdentify(l.ctx, &apply.CheckIdentifyReq{
		UserId: tempId,
	})
	plaintiffId := plaintiff.ScholarId //申诉学者id
	defendantUser, _ := l.svcCtx.ApplyRpc.CheckUser(l.ctx, &apply.CheckUserReq{ScholarId: defendantId})
	defendantUserId := defendantUser.UserId //被申诉用户的id
	newGrievance := model.Grievance{
		PlaintiffId: plaintiffId,
		DefendantId: defendantId,
		PaperId:     paperId,
	}
	tempGrievance, err := l.svcCtx.GrievanceModel.Insert(l.ctx, &newGrievance)
	gId, _ := tempGrievance.LastInsertId()
	paper, _ := l.svcCtx.PaperRpc.GetPaper(l.ctx, &paper2.GetPaperReq{PaperId: paperId})
	paperName := paper.PaperName
	content := fmt.Sprintf("%s 对你的文献 %s 发起申诉", username, paperName)
	//给被申诉者发通知
	_, _ = l.svcCtx.MessageRpc.CreateMessage(l.ctx, &message2.CreateMessageReq{
		ReceiverId:  defendantUserId,
		Content:     content,
		MessageType: 4,
		UId:         tempId,
		GId:         gId,
		PId:         paperId,
	})
	return &types.LaunchGrievanceResponse{}, nil
}
