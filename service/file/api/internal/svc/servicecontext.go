package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"soft2_backend/service/file/api/internal/config"
	"soft2_backend/service/file/model"
)

type ServiceContext struct {
	Config          config.Config
	ApplyFileModel  model.ApplyFileModel
	HelpFileModel   model.HelpFileModel
	UserAvatarModel model.UserAvatarModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:          c,
		ApplyFileModel:  model.NewApplyFileModel(conn),
		UserAvatarModel: model.NewUserAvatarModel(conn),
		HelpFileModel:   model.NewHelpFileModel(conn),
	}
}
