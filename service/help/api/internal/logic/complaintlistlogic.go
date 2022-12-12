package logic

import (
	"context"
	"soft2_backend/service/help/api/internal/svc"
	"soft2_backend/service/help/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ComplaintListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewComplaintListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ComplaintListLogic {
	return &ComplaintListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ComplaintListLogic) ComplaintList(req *types.ComplaintListReq) (resp *types.ComplaintListReply, err error) {
	//reqList, err := l.svcCtx.LiteratureRequestModel.FindComplaint(l.ctx)
	//sum := len(reqList)
	//var reql []types.Complaint
	//for i, oneReq := range reqList {
	//	if req.End != -1 && i >= int(req.End) {
	//		break
	//	}
	//	if i >= int(req.Start) {
	//		var request types.Complaint
	//		request.RequestId = oneReq.Id
	//		request.RequestTitle = oneReq.Title
	//		request.RequestTime = oneReq.RequestTime.Format("2006-1-02 15:04")
	//		request.Content = oneReq.Complaint
	//		helpFile, _ := l.svcCtx.FileRpc.GetHelpFile(l.ctx, &file.ApplyIdReq{
	//			Id: oneReq.Id,
	//		})
	//		request.Url = helpFile.Url
	//		reql = append(reql, request)
	//	}
	//}
	//return &types.ComplaintListReply{
	//	Complaints: reql,
	//	Num:        int64(sum),
	//}, nil
	return nil, nil
}
