package logic

import (
	"context"
	"fmt"
	"soft2_backend/service/apply/rpc/types/apply"
	message2 "soft2_backend/service/message/rpc/types/message"
	paper2 "soft2_backend/service/paper/rpc/paper"
	"soft2_backend/service/paper/rpc/streamgreeter"
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
	plaintiffId := grievance.PlaintiffId //申诉学者id
	defendantId := grievance.DefendantId //被申诉学者id
	paperId := grievance.PaperId
	plaintiff, _ := l.svcCtx.ApplyRpc.CheckUser(l.ctx, &apply.CheckUserReq{ScholarId: plaintiffId})
	paper, _ := l.svcCtx.PaperRpc.GetPaper(l.ctx, &paper2.GetPaperReq{PaperId: paperId})
	_, _ = l.svcCtx.PaperRpc.MovePaper(l.ctx, &streamgreeter.MovePaperReq{
		PaperId:  paperId,
		OwnerId:  defendantId,
		TargetId: plaintiffId,
	})
	//告知申诉结果
	content := fmt.Sprintf("你对文献%s 的申诉通过", paper.PaperName)
	_, _ = l.svcCtx.MessageRpc.CreateMessage(l.ctx, &message2.CreateMessageReq{
		ReceiverId:  plaintiff.UserId,
		Content:     content,
		MessageType: 6,
		Result:      0,
		GId:         req.GrievanceId,
		PId:         paperId,
	})
	return &types.GrievanceAcceptResponse{}, nil
}
