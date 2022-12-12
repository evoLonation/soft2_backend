package logic

import (
	"context"
	"database/sql"
	"soft2_backend/service/message/model"
	"soft2_backend/service/message/rpc/internal/svc"
	"soft2_backend/service/message/rpc/types/message"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMessageLogic {
	return &CreateMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateMessageLogic) CreateMessage(in *message.CreateMessageReq) (*message.CreateMessageReply, error) {

	newMessage := model.Message{
		ReceiverId:  in.ReceiverId,
		Content:     in.Content,
		MessageType: in.MessageType,
		//Result:      sql.NullInt64{},
		//UId:         sql.NullInt64{},
		//GId:         sql.NullInt64{},
		//PId:         sql.NullInt64{},
		//RId:         sql.NullInt64{},
	}
	switch in.MessageType {
	case 1:
		newMessage.UId = sql.NullInt64{
			Int64: in.UId,
			Valid: true,
		}
		newMessage.PId = sql.NullString{
			String: in.PId,
			Valid:  true,
		}
		break
	case 2:
		newMessage.UId = sql.NullInt64{
			Int64: in.UId,
			Valid: true,
		}
		newMessage.PId = sql.NullString{
			String: in.PId,
			Valid:  true,
		}
		break
	case 3:
		newMessage.UId = sql.NullInt64{
			Int64: in.UId,
			Valid: true,
		}
		break
	case 4:
		newMessage.UId = sql.NullInt64{
			Int64: in.UId,
			Valid: true,
		}
		newMessage.GId = sql.NullInt64{
			Int64: in.GId,
			Valid: true,
		}
		newMessage.PId = sql.NullString{
			String: in.PId,
			Valid:  true,
		}
		break
	case 5:
		newMessage.Result = sql.NullInt64{
			Int64: in.Result,
			Valid: true,
		}
		break
	case 6:
		newMessage.GId = sql.NullInt64{
			Int64: in.GId,
			Valid: true,
		}
		newMessage.PId = sql.NullString{
			String: in.PId,
			Valid:  true,
		}
		newMessage.Result = sql.NullInt64{
			Int64: in.Result,
			Valid: true,
		}
		break
	case 7:
		newMessage.RId = sql.NullInt64{
			Int64: in.RId,
			Valid: true,
		}
		break
	case 8:
		newMessage.RId = sql.NullInt64{
			Int64: in.RId,
			Valid: true,
		}
		break
	}

	_, err := l.svcCtx.MessageModel.Insert(l.ctx, &newMessage)
	if err != nil {
		return nil, err
	}

	return &message.CreateMessageReply{}, nil
}
