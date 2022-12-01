package logic

import (
	"context"
	message2 "soft2_backend/service/message/rpc/types/message"
	"soft2_backend/service/paper/rpc/streamgreeter"
	"strconv"

	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GrievanceAcceptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGrievanceAcceptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GrievanceAcceptLogic {
	return &GrievanceAcceptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GrievanceAcceptLogic) GrievanceAccept(req *types.GrievanceAcceptRequest) (resp *types.GrievanceAcceptResponse, err error) {
	//移动学者认领的文献
	grievance, err := l.svcCtx.GrievanceModel.FindOne(l.ctx, req.GrievanceId)
	plaintiffId := grievance.PlaintiffId
	defendantId := grievance.DefendantId
	paperId := grievance.PaperId
	_, _ = l.svcCtx.PaperRpc.MovePaper(l.ctx, &streamgreeter.MovePaperReq{
		PaperId:  strconv.FormatInt(paperId, 10),
		OwnerId:  strconv.FormatInt(defendantId, 10),
		TargetId: strconv.FormatInt(plaintiffId, 10),
	})
	//告知申诉结果
	_, _ = l.svcCtx.MessageRpc.CreateMessage(l.ctx, &message2.CreateMessageReq{
		ReceiverId:  plaintiffId,
		Content:     "",
		MessageType: 6,
		Result:      0,
		GId:         req.GrievanceId,
		PId:         strconv.FormatInt(paperId, 10),
	})
	return &types.GrievanceAcceptResponse{}, nil
}
