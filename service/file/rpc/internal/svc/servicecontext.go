package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"soft2_backend/service/apply/rpc/applyclient"
	"soft2_backend/service/file/model"
	"soft2_backend/service/file/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	model.UserAvatarModel
	model.HelpFileModel
	applyclient.Apply
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:          c,
		UserAvatarModel: model.NewUserAvatarModel(conn),
		HelpFileModel:   model.NewHelpFileModel(conn),
		Apply:           applyclient.NewApply(zrpc.MustNewClient(c.ApplyRpcConf)),
	}
}
