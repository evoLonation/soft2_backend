package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-share/message/api/internal/config"
	"go-zero-share/message/model"
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
