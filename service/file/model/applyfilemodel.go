package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ApplyFileModel = (*customApplyFileModel)(nil)

type (
	// ApplyFileModel is an interface to be customized, add more methods here,
	// and implement the added methods in customApplyFileModel.
	ApplyFileModel interface {
		applyFileModel
	}

	customApplyFileModel struct {
		*defaultApplyFileModel
	}
)

// NewApplyFileModel returns a model for the database table.
func NewApplyFileModel(conn sqlx.SqlConn) ApplyFileModel {
	return &customApplyFileModel{
		defaultApplyFileModel: newApplyFileModel(conn),
	}
}
