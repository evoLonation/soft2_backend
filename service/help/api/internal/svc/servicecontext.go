package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"soft2_backend/service/file/rpc/fileclient"
	"soft2_backend/service/help/api/internal/config"
	"soft2_backend/service/help/model"
	"soft2_backend/service/message/rpc/messageclient"
)

type ServiceContext struct {
	Config                 config.Config
	LiteratureRequestModel model.LiteratureRequestModel
	UserHelpModel          model.UserHelpModel
	LiteratureHelpModel    model.LiteratureHelpModel
	FileRpc                fileclient.File
	MessageRpc             messageclient.Message
	Auth                   struct {
		AccessSecret string
		AccessExpire int64
	}
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:                 c,
		LiteratureRequestModel: model.NewLiteratureRequestModel(conn),
		UserHelpModel:          model.NewUserHelpModel(conn),
		LiteratureHelpModel:    model.NewLiteratureHelpModel(conn),
		FileRpc:                fileclient.NewFile(zrpc.MustNewClient(c.FireRpc)),
		MessageRpc:             messageclient.NewMessage(zrpc.MustNewClient(c.MessageRpc)),
	}
}
