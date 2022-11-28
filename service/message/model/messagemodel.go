package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MessageModel = (*customMessageModel)(nil)

type (
	// MessageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageModel.
	MessageModel interface {
		messageModel

		FindAllByUser(ctx context.Context, userId int64) ([]*Message, error)
	}

	customMessageModel struct {
		*defaultMessageModel
	}
)

func (m *customMessageModel) FindAllByUser(ctx context.Context, userId int64) ([]*Message, error) {
	var resp []*Message

	query := fmt.Sprintf("select %s from %s where `userId` = ?", messageRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewMessageModel returns a model for the database table.
func NewMessageModel(conn sqlx.SqlConn) MessageModel {
	return &customMessageModel{
		defaultMessageModel: newMessageModel(conn),
	}
}
