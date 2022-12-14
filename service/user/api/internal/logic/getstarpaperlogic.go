package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"soft2_backend/service/paper/rpc/streamgreeter"
	"soft2_backend/service/user/api/internal/svc"
	"soft2_backend/service/user/api/internal/types"
)

type GetStarPaperLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStarPaperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStarPaperLogic {
	return &GetStarPaperLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStarPaperLogic) GetStarPaper(req *types.GetStarPaperRequest) (resp *types.GetStarPaperResponse, err error) {
	// todo: add your logic here and delete this line
	userId, _ := l.ctx.Value("UserId").(json.Number).Int64()
	reqList, err := l.svcCtx.CollectModel.FindByUserId(l.ctx, userId) //获取收藏的文献
	sum := len(reqList)
	if sum == 0 {
		return &types.GetStarPaperResponse{PaperStar: nil}, nil
	}
	var paperIds []string
	for i := 0; i < sum; i++ {
		paperIds = append(paperIds, reqList[i].PaperId)
	} //获取收藏的文献id
	ListPaperReply, err := l.svcCtx.PaperRpc.ListGetPaper(l.ctx, &streamgreeter.ListGetPaperReq{PaperId: paperIds}) //获取收藏的文献详情
	var reql = make([]types.PaperStarReply, sum)
	for i := 0; i < sum; i++ {
		reql[i].PaperId = reqList[i].PaperId
		reql[i].PaperName = ListPaperReply.Papers[i].PaperName
		reql[i].Org = ListPaperReply.Papers[i].Org
		reql[i].Date = ListPaperReply.Papers[i].Year
		Authors := ListPaperReply.Papers[i].Authors
		var authorReply []types.AuthorReply
		authsum := len(Authors)
		for i2, author := range Authors {
			if i2 > authsum {
				break
			}
			var temp types.AuthorReply
			temp.Id = author.Id
			temp.Name = author.Name
			authorReply = append(authorReply, temp)
		}
		reql[i].Author = authorReply
	}
	return &types.GetStarPaperResponse{PaperStar: reql}, nil
}
