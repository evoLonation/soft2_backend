package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"soft2_backend/service/apply/api/internal/svc"
	"soft2_backend/service/apply/api/internal/types"
	"soft2_backend/service/paper/rpc/streamgreeter"
)

type GetApplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApplyLogic {
	return &GetApplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApplyLogic) GetApply(req *types.GetApplyRequest) (resp *types.GetApplyResponse, err error) {
	list, err := l.svcCtx.ApplyModel.FindAll(l.ctx)

	infoList := make([]types.ApplyInfo, 0)

	i := req.Start
	end := int64(len(list))
	if req.End != 0 && req.End < end {
		end = req.End
	}
	count := end - i

	checkScholars := make([]string, 0)
	for ; i < end; i++ {
		checkScholars = append(checkScholars, list[i].ScholarId)
	}
	scholars, err := l.svcCtx.PaperRpc.ListCheckScholar(l.ctx, &streamgreeter.ListCheckScholarReq{ScholarId: checkScholars})
	if err != nil {
		return nil, err
	}

	for i := req.Start; i < end; i++ {

		info := types.ApplyInfo{
			ApplyId:     list[i].ApplyId,
			ScholarName: scholars.Scholars[i-req.Start].ScholarName,
			Institution: scholars.Scholars[i-req.Start].Org,
			ApplyType:   list[i].ApplyType,
			ApplyTime:   list[i].ApplyTime.Format("2006-01-02 15:04:05"),
		}

		if list[i].Email.Valid {
			info.Email = list[i].Email.String
		} else if list[i].Url.Valid {
			info.URL = list[i].Url.String
		}

		infoList = append(infoList, info)
	}

	return &types.GetApplyResponse{
		Count:     count,
		ApplyList: infoList,
	}, nil
}
