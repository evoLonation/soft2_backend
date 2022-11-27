package model

import (
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CollectModel = (*customCollectModel)(nil)

type (
	// CollectModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCollectModel.
	CollectModel interface {
		collectModel
	}

	customCollectModel struct {
		*defaultCollectModel
	}
)

// NewCollectModel returns a model for the database table.
func NewCollectModel(conn sqlx.SqlConn, c cache.CacheConf) CollectModel {
	return &customCollectModel{
		defaultCollectModel: newCollectModel(conn, c),
	}
}
