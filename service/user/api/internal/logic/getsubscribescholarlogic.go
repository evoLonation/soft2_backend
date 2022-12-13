package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"soft2_backend/service/paper/rpc/streamgreeter"
	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"
)

type GetSubscribeScholarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSubscribeScholarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubscribeScholarLogic {
	return &GetSubscribeScholarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSubscribeScholarLogic) GetSubscribeScholar(req *types.GetSubscribeScholarRequest) (resp *types.GetSubscribeScholarResponse, err error) {
	// todo: add your logic here and delete this line
	userId, _ := l.ctx.Value("UserId").(json.Number).Int64()
	reqList, err := l.svcCtx.SubscribeModel.FindByUserId(l.ctx, userId)
	sum := len(reqList)
	if sum == 0 {
		return &types.GetSubscribeScholarResponse{ScholarSubscribe: nil}, nil
	}
	var scholarIds []string
	for i := 0; i < sum; i++ {
		scholarIds = append(scholarIds, reqList[i].ScholarId)
	}
	ListScholarReply, err := l.svcCtx.PaperRpc.ListCheckScholar(l.ctx, &streamgreeter.ListCheckScholarReq{ScholarId: scholarIds})
	var reql = make([]types.ScholarSubscribeReply, sum)
	for i := 0; i < sum; i++ {
		reql[i].ScholarId = reqList[i].ScholarId
		reql[i].ScholarName = ListScholarReply.Scholars[i].ScholarName
		reql[i].Org = ListScholarReply.Scholars[i].Org
		reql[i].PaperNum = ListScholarReply.Scholars[i].PaperNum
		reql[i].Url = ListScholarReply.Scholars[i].Url
	}
	return &types.GetSubscribeScholarResponse{ScholarSubscribe: reql}, nil
}
