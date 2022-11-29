package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"soft2_backend/service/file/model"
	"soft2_backend/service/file/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	model.UserAvatarModel
	model.HelpFileModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:          c,
		UserAvatarModel: model.NewUserAvatarModel(conn),
		HelpFileModel:   model.NewHelpFileModel(conn),
	}
}
