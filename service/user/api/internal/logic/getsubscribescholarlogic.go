package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"soft2_backend/service/paper/rpc/streamgreeter"
	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"
	"soft2_backend/service/user/model"

	"github.com/zeromicro/go-zero/core/logx"
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
	if err == model.ErrNotFound {
		return &types.GetSubscribeScholarResponse{ScholarSubscribe: nil}, nil
	}
	fmt.Printf("11111111111111\n%d\n", len(reqList))
	sum := len(reqList)
	var scholarIds []string
	for i := 0; i < sum; i++ {
		scholarIds = append(scholarIds, reqList[i].ScholarId)
	}
	fmt.Printf("222222222222\n%s", scholarIds[0])
	ListScholarReply, err := l.svcCtx.PaperRpc.ListCheckScholar(l.ctx, &streamgreeter.ListCheckScholarReq{ScholarId: scholarIds})
	fmt.Printf("33333333333\n%s", ListScholarReply.Scholars[0].ScholarName)
	var reql []types.ScholarSubscribeReply
	for i := 0; i < sum; i++ {
		fmt.Printf("4444444444444\n%d\n", i)
		reql[i].ScholarId = reqList[i].ScholarId
		fmt.Printf("55555555555555\n")
		reql[i].ScholarName = ListScholarReply.Scholars[i].ScholarName
		fmt.Printf("6666666666666\n")
		reql[i].Org = ListScholarReply.Scholars[i].Org
		fmt.Printf("77777777777777\n")
		reql[i].PaperNum = ListScholarReply.Scholars[i].PaperNum
		fmt.Printf("8888888888888888\n")
		reql[i].Url = ListScholarReply.Scholars[i].Url
	}
	return &types.GetSubscribeScholarResponse{ScholarSubscribe: reql}, nil
}
