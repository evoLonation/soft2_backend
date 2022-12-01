package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"soft2_backend/service/apply/rpc/applyclient"
	"soft2_backend/service/help/rpc/helpclient"
	"soft2_backend/service/message/rpc/messageclient"
	"soft2_backend/service/paper/rpc/streamgreeter"
	"soft2_backend/service/user/api/internal/config"
	"soft2_backend/service/user/model"
)

type ServiceContext struct {
	Config         config.Config
	UserModel      model.UserModel
	CollectModel   model.CollectModel
	CommentModel   model.CommentModel
	LikeModel      model.LikeModel
	SubscribeModel model.SubscribeModel
	GrievanceModel model.GrievanceModel
	ApplyRpc       applyclient.Apply
	MessageRpc     messageclient.Message
	PaperRpc       streamgreeter.StreamGreeter
	HelpRpc        helpclient.Help
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:         c,
		UserModel:      model.NewUserModel(conn),
		CollectModel:   model.NewCollectModel(conn),
		CommentModel:   model.NewCommentModel(conn),
		LikeModel:      model.NewLikeModel(conn),
		SubscribeModel: model.NewSubscribeModel(conn),
		GrievanceModel: model.NewGrievanceModel(conn),
		ApplyRpc:       applyclient.NewApply(zrpc.MustNewClient(c.ApplyRpc)),
		MessageRpc:     messageclient.NewMessage(zrpc.MustNewClient(c.MessageRpc)),
		PaperRpc:       streamgreeter.NewStreamGreeter(zrpc.MustNewClient(c.PaperRpc)),
		HelpRpc:        helpclient.NewHelp(zrpc.MustNewClient(c.HelpRpc)),
	}
}
