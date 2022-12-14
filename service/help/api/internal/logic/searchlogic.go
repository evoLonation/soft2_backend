package logic

import (
	"context"
	"soft2_backend/service/help/api/internal/svc"
	"soft2_backend/service/help/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.ReqSearchReq) (resp *types.ReqSearchReply, err error) {
	// todo: add your logic here and delete this line
	search := "'" + req.SearchContent + "'"
	reqList, err := l.svcCtx.LiteratureRequestModel.FindByContent(l.ctx, search)
	sum := len(reqList)
	var reql []types.Search
	for i, oneReq := range reqList {
		if req.End != -1 && i >= int(req.End) {
			break
		}
		if i >= int(req.Start) {
			var request types.Search
			request.ReqTitle = oneReq.Title
			request.RequestId = oneReq.Id
			request.RequestTime = oneReq.RequestTime.Format("2006-1-02 15:04")
			request.RequestContent = oneReq.RequestContent
			request.Wealth = oneReq.Wealth
			request.RequestStatus = oneReq.RequestStatus
			request.Author = oneReq.Author
			request.Link = oneReq.Link
			request.Magazine = oneReq.Magazine
			reql = append(reql, request)
		}
	}
	return &types.ReqSearchReply{
		ReqList:     reql,
		RequestsNum: int64(sum),
	}, nil
}
