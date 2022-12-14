package logic

import (
	"context"
	"encoding/json"
	"soft2_backend/service/message/api/internal/svc"
	"soft2_backend/service/message/api/internal/types"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageLogic {
	return &GetMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMessageLogic) GetMessage(req *types.GetMessageRequest) (resp *types.GetMessageResponse, err error) {
	userId, _ := l.ctx.Value("UserId").(json.Number).Int64()
	list, err := l.svcCtx.MessageModel.FindAllByUser(l.ctx, userId)

	infoList := make([]types.MessageInfo, 0)

	i := req.Start
	end := int64(len(list))
	if req.End != 0 && req.End < end {
		end = req.End
	}
	count := end - i
	var cstZone = time.FixedZone("CST", 8*3600)
	for ; i < end; i++ {
		info := types.MessageInfo{
			MessageId:   list[i].MsgId,
			MessageType: list[i].MessageType,
			Content:     list[i].Content,
			Read:        list[i].Read,
			MessageTime: list[i].MsgTime.In(cstZone).Format("2006-01-02 15:04:05"),
		}

		if list[i].UId.Valid {
			info.UId = list[i].UId.Int64
		}
		if list[i].SId.Valid {
			info.SId = list[i].SId.String
		}
		if list[i].GId.Valid {
			info.GId = list[i].GId.Int64
		}
		if list[i].PId.Valid {
			info.PId = list[i].PId.String
		}
		if list[i].RId.Valid {
			info.RId = list[i].RId.Int64
		}
		if list[i].Result.Valid {
			info.Result = list[i].Result.Int64
		}

		infoList = append(infoList, info)
	}

	return &types.GetMessageResponse{
		Count:       count,
		MessageList: infoList,
	}, nil
}
