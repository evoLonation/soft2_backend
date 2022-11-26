package model

import (
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserAvatarModel = (*customUserAvatarModel)(nil)

type (
	// UserAvatarModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserAvatarModel.
	UserAvatarModel interface {
		userAvatarModel
	}

	customUserAvatarModel struct {
		*defaultUserAvatarModel
	}
)

// NewUserAvatarModel returns a model for the database table.
func NewUserAvatarModel(conn sqlx.SqlConn) UserAvatarModel {
	return &customUserAvatarModel{
		defaultUserAvatarModel: newUserAvatarModel(conn),
	}
}
