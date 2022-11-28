package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"soft2_backend/service/help/api/internal/config"
	"soft2_backend/service/help/model"
)

type ServiceContext struct {
	Config                 config.Config
	LiteratureRequestModel model.LiteratureRequestModel
	UserHelpModel          model.UserHelpModel
	LiteratureHelpModel    model.LiteratureHelpModel
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
	}
}
