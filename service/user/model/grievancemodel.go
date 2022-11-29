package model

import (
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GrievanceModel = (*customGrievanceModel)(nil)

type (
	// GrievanceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGrievanceModel.
	GrievanceModel interface {
		grievanceModel
	}

	customGrievanceModel struct {
		*defaultGrievanceModel
	}
)

// NewGrievanceModel returns a model for the database table.
func NewGrievanceModel(conn sqlx.SqlConn, c cache.CacheConf) GrievanceModel {
	return &customGrievanceModel{
		defaultGrievanceModel: newGrievanceModel(conn, c),
	}
}
