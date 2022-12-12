package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"soft2_backend/service/help/model"
	"soft2_backend/service/help/rpc/internal/config"
	"soft2_backend/service/message/rpc/messageclient"
)

type ServiceContext struct {
	Config                 config.Config
	LiteratureHelpModel    model.LiteratureHelpModel
	LiteratureRequestModel model.LiteratureRequestModel
	UserHelpModel          model.UserHelpModel
	MessageRpc             messageclient.Message
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:                 c,
		LiteratureRequestModel: model.NewLiteratureRequestModel(conn),
		LiteratureHelpModel:    model.NewLiteratureHelpModel(conn),
		UserHelpModel:          model.NewUserHelpModel(conn),
		MessageRpc:             messageclient.NewMessage(zrpc.MustNewClient(c.MessageRpc)),
	}
}
