package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"soft2_backend/service/apply/api/internal/config"
	"soft2_backend/service/apply/model"
	"soft2_backend/service/paper/rpc/streamgreeter"
)

type ServiceContext struct {
	Config          config.Config
	ApplyModel      model.ApplyModel
	VerifycodeModel model.VerifycodeModel

	PaperRpc streamgreeter.StreamGreeter
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:          c,
		ApplyModel:      model.NewApplyModel(conn),
		VerifycodeModel: model.NewVerifycodeModel(conn),
		PaperRpc:        streamgreeter.NewStreamGreeter(zrpc.MustNewClient(c.PaperRpc)),
	}
}
