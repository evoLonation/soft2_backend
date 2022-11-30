package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"soft2_backend/service/apply/rpc/applyclient"
	"soft2_backend/service/message/rpc/messageclient"
	"soft2_backend/service/paper/rpc/streamgreeter"
	"soft2_backend/service/user/api/internal/config"
	"soft2_backend/service/user/model"
)

type ServiceContext struct {
	Config         config.Config
	UserModel      model.UserModel
	CollectModel   model.CollectModel
	CommentModel   model.CommentModel
	LikeModel      model.LikeModel
	SubscribeModel model.SubscribeModel
	GrievanceModel model.GrievanceModel
	ApplyRpc       applyclient.Apply
	MessageRpc     messageclient.Message
	PaperRpc       streamgreeter.StreamGreeter
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:         c,
		UserModel:      model.NewUserModel(conn, c.CacheRedis),
		CollectModel:   model.NewCollectModel(conn, c.CacheRedis),
		CommentModel:   model.NewCommentModel(conn, c.CacheRedis),
		LikeModel:      model.NewLikeModel(conn, c.CacheRedis),
		SubscribeModel: model.NewSubscribeModel(conn, c.CacheRedis),
		GrievanceModel: model.NewGrievanceModel(conn, c.CacheRedis),
		ApplyRpc:       applyclient.NewApply(zrpc.MustNewClient(c.ApplyRpc)),
		MessageRpc:     messageclient.NewMessage(zrpc.MustNewClient(c.MessageRpc)),
		PaperRpc:       streamgreeter.NewStreamGreeter(zrpc.MustNewClient(c.PaperRpc)),
	}
}
