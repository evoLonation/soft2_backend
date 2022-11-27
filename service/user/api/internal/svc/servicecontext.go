package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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
	}
}
