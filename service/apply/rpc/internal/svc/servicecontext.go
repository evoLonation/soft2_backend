package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"soft2_backend/service/apply/model"
	"soft2_backend/service/apply/rpc/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	ApplyModel model.ApplyModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		ApplyModel: model.NewApplyModel(conn),
	}
}
