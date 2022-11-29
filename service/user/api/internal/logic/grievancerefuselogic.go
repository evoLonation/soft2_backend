package logic

import (
	"context"
	message2 "soft2_backend/service/message/rpc/types/message"

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
	paperId := grievance.PaperId
	_, _ = l.svcCtx.MessageRpc.CreateMessage(l.ctx, &message2.CreateMessageReq{
		ReceiverId:  plantiffId,
		Content:     "",
		MessageType: 6,
		Result:      1,
		GId:         req.GrievanceId,
		PId:         paperId,
	})
	return &types.GrievanceRefuseResponse{}, nil
}
