package model

import (
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserHelpModel = (*customUserHelpModel)(nil)

type (
	// UserHelpModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserHelpModel.
	UserHelpModel interface {
		userHelpModel
	}

	customUserHelpModel struct {
		*defaultUserHelpModel
	}
)

// NewUserHelpModel returns a model for the database table.
func NewUserHelpModel(conn sqlx.SqlConn) UserHelpModel {
	return &customUserHelpModel{
		defaultUserHelpModel: newUserHelpModel(conn),
	}
}
