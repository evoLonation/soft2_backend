package model

import (
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LiteratureRequestModel = (*customLiteratureRequestModel)(nil)

type (
	// LiteratureRequestModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLiteratureRequestModel.
	LiteratureRequestModel interface {
		literatureRequestModel
	}

	customLiteratureRequestModel struct {
		*defaultLiteratureRequestModel
	}
)

// NewLiteratureRequestModel returns a model for the database table.
func NewLiteratureRequestModel(conn sqlx.SqlConn) LiteratureRequestModel {
	return &customLiteratureRequestModel{
		defaultLiteratureRequestModel: newLiteratureRequestModel(conn),
	}
}
