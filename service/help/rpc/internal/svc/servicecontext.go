package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"soft2_backend/service/help/model"
	"soft2_backend/service/help/rpc/internal/config"
)

type ServiceContext struct {
	Config                 config.Config
	LiteratureHelpModel    model.LiteratureHelpModel
	LiteratureRequestModel model.LiteratureRequestModel
	UserHelpModel          model.UserHelpModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:                 c,
		LiteratureRequestModel: model.NewLiteratureRequestModel(conn),
		LiteratureHelpModel:    model.NewLiteratureHelpModel(conn),
		UserHelpModel:          model.NewUserHelpModel(conn),
	}
}
