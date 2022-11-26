package model

import (
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HelpFileModel = (*customHelpFileModel)(nil)

type (
	// HelpFileModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHelpFileModel.
	HelpFileModel interface {
		helpFileModel
	}

	customHelpFileModel struct {
		*defaultHelpFileModel
	}
)

// NewHelpFileModel returns a model for the database table.
func NewHelpFileModel(conn sqlx.SqlConn) HelpFileModel {
	return &customHelpFileModel{
		defaultHelpFileModel: newHelpFileModel(conn),
	}
}
