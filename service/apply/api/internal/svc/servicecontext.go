package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-share/apply/api/internal/config"
	"go-zero-share/apply/model"
)

type ServiceContext struct {
	Config     config.Config
	ApplyModel model.ApplyModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		ApplyModel: model.NewApplyModel(conn, c.CacheRedis),
	}
}
