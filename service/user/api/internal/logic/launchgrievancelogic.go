package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"soft2_backend/service/apply/rpc/types/apply"
	message2 "soft2_backend/service/message/rpc/types/message"
	paper2 "soft2_backend/service/paper/rpc/paper"
	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"
	"soft2_backend/service/user/model"
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
	paperId := req.PaperId
	defendantScholarId := req.ScholarId
	plaintiffId, _ := l.ctx.Value("UserId").(json.Number).Int64()                                               //申诉用户id
	plaintiffUser, _ := l.svcCtx.UserModel.FindOne(l.ctx, plaintiffId)                                          //申诉用户
	plaintiffScholar, _ := l.svcCtx.ApplyRpc.CheckIdentify(l.ctx, &apply.CheckIdentifyReq{UserId: plaintiffId}) //申诉学者
	plaintiffScholarId := plaintiffScholar.ScholarId                                                            //申诉学者id
	defendantScholar, _ := l.svcCtx.ApplyRpc.CheckUser(l.ctx, &apply.CheckUserReq{ScholarId: defendantScholarId})
	defendantScholarUserId := defendantScholar.UserId //被申诉者用户id
	newGrievance := model.Grievance{
		PlaintiffId: plaintiffScholarId,
		DefendantId: req.ScholarId,
		PaperId:     paperId,
	}
	tempGrievance, _ := l.svcCtx.GrievanceModel.Insert(l.ctx, &newGrievance)
	gId, _ := tempGrievance.LastInsertId()
	paper, _ := l.svcCtx.PaperRpc.GetPaper(l.ctx, &paper2.GetPaperReq{PaperId: paperId})
	var username string
	var papername string
	if len(plaintiffUser.Nickname) > 20 {
		username = plaintiffUser.Nickname[0:20] + "..."
	} else {
		username = plaintiffUser.Nickname
	}

	if len(paper.PaperName) > 20 {
		papername = paper.PaperName[0:20] + "..."
	} else {
		papername = paper.PaperName
	}
	content := fmt.Sprintf("%s对你的文献%s发起申诉", username, papername)
	fmt.Printf("1111111111111\n%d\n", defendantScholarUserId)
	fmt.Printf("2222222222222\n%s\n", plaintiffScholarId)
	_, err = l.svcCtx.MessageRpc.CreateMessage(l.ctx, &message2.CreateMessageReq{
		ReceiverId:  defendantScholarUserId,
		Content:     content,
		MessageType: 4,
		SId:         plaintiffScholarId,
		GId:         gId,
		PId:         paperId,
	})
	if err != nil {
		fmt.Printf("%s", err)
	}
	return &types.LaunchGrievanceResponse{}, nil
}
