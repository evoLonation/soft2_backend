package model

import (
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LiteratureHelpModel = (*customLiteratureHelpModel)(nil)

type (
	// LiteratureHelpModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLiteratureHelpModel.
	LiteratureHelpModel interface {
		literatureHelpModel
	}

	customLiteratureHelpModel struct {
		*defaultLiteratureHelpModel
	}
)

// NewLiteratureHelpModel returns a model for the database table.
func NewLiteratureHelpModel(conn sqlx.SqlConn) LiteratureHelpModel {
	return &customLiteratureHelpModel{
		defaultLiteratureHelpModel: newLiteratureHelpModel(conn),
	}
}
