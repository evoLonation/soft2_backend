package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"soft2_backend/service/apply/rpc/applyclient"
	"soft2_backend/service/file/api/internal/config"
	"soft2_backend/service/file/model"
	"soft2_backend/service/help/rpc/helpclient"
	"soft2_backend/service/message/rpc/messageclient"
)

type ServiceContext struct {
	config.Config
	model.HelpFileModel
	model.UserAvatarModel
	helpclient.Help
	applyclient.Apply
	messageclient.Message
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:          c,
		UserAvatarModel: model.NewUserAvatarModel(conn),
		HelpFileModel:   model.NewHelpFileModel(conn),
		Help:            helpclient.NewHelp(zrpc.MustNewClient(c.HelpRpcConf)),
		Apply:           applyclient.NewApply(zrpc.MustNewClient(c.ApplyRpcConf)),
		Message:         messageclient.NewMessage(zrpc.MustNewClient(c.MessageRpcConf)),
	}
}
