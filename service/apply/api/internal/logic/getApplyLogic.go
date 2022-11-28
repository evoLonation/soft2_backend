package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-share/apply/api/internal/svc"
	"go-zero-share/apply/api/internal/types"
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

	for ; i < end; i++ {
		// todo 调用学者rpc查找scholarname和institution

		info := types.ApplyInfo{
			ApplyId:     list[i].ApplyId,
			ScholarName: "lll",
			Institution: "www",
			ApplyType:   list[i].ApplyType,
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
