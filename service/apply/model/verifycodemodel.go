package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ VerifycodeModel = (*customVerifycodeModel)(nil)

type (
	// VerifycodeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVerifycodeModel.
	VerifycodeModel interface {
		verifycodeModel
	}

	customVerifycodeModel struct {
		*defaultVerifycodeModel
	}
)

// NewVerifycodeModel returns a model for the database table.
func NewVerifycodeModel(conn sqlx.SqlConn) VerifycodeModel {
	return &customVerifycodeModel{
		defaultVerifycodeModel: newVerifycodeModel(conn),
	}
}
