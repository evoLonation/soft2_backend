package logic

import (
	"context"
	"encoding/json"
	"soft2_backend/service/apply/rpc/types/apply"
	message2 "soft2_backend/service/message/rpc/types/message"
	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"
	"soft2_backend/service/user/model"
	"strconv"

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
	plaintiff, _ := l.svcCtx.ApplyRpc.CheckIdentify(l.ctx, &apply.CheckIdentifyReq{
		UserId: tempId,
	})
	plaintiffId, err := strconv.ParseInt(plaintiff.ScholarId, 10, 64)
	newGrievance := model.Grievance{
		PlaintiffId: plaintiffId,
		DefendantId: defendantId,
		PaperId:     paperId,
	}
	tempGrievance, err := l.svcCtx.GrievanceModel.Insert(l.ctx, &newGrievance)
	gId, _ := tempGrievance.LastInsertId()
	//给被申诉者发通知
	_, _ = l.svcCtx.MessageRpc.CreateMessage(l.ctx, &message2.CreateMessageReq{
		ReceiverId:  defendantId,
		Content:     "",
		MessageType: 4,
		UId:         plaintiffId,
		GId:         gId,
		PId:         strconv.FormatInt(paperId, 10),
	})
	return &types.LaunchGrievanceResponse{}, nil
}
