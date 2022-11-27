package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SubscribeModel = (*customSubscribeModel)(nil)

type (
	// SubscribeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSubscribeModel.
	SubscribeModel interface {
		subscribeModel
	}

	customSubscribeModel struct {
		*defaultSubscribeModel
	}
)

// NewSubscribeModel returns a model for the database table.
func NewSubscribeModel(conn sqlx.SqlConn, c cache.CacheConf) SubscribeModel {
	return &customSubscribeModel{
		defaultSubscribeModel: newSubscribeModel(conn, c),
	}
}
