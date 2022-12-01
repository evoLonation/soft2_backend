package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"soft2_backend/service/message/model"
	"soft2_backend/service/message/rpc/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	MessageModel model.MessageModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		MessageModel: model.NewMessageModel(conn),
	}
}
