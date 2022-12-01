package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ApplyModel = (*customApplyModel)(nil)

type (
	// ApplyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customApplyModel.
	ApplyModel interface {
		applyModel

		FindAll(ctx context.Context) ([]*Apply, error)
		FindByUserId(ctx context.Context, userId int64) (*Apply, error)
		FindByScholarId(ctx context.Context, scholarId string) (*Apply, error)
	}

	customApplyModel struct {
		*defaultApplyModel
	}
)

func (m *customApplyModel) FindAll(ctx context.Context) ([]*Apply, error) {
	var resp []*Apply

	query := fmt.Sprintf("select %s from %s where `status` = 0", applyRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customApplyModel) FindByUserId(ctx context.Context, userId int64) (*Apply, error) {
	query := fmt.Sprintf("select %s from %s where `userId` = ? and `status` = 1", applyRows, m.table)
	var resp Apply
	err := m.conn.QueryRowCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customApplyModel) FindByScholarId(ctx context.Context, scholarId string) (*Apply, error) {
	query := fmt.Sprintf("select %s from %s where `scholarId` = ? and `status` = 1", applyRows, m.table)
	var resp Apply
	err := m.conn.QueryRowCtx(ctx, &resp, query, scholarId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewApplyModel returns a model for the database table.
func NewApplyModel(conn sqlx.SqlConn) ApplyModel {
	return &customApplyModel{
		defaultApplyModel: newApplyModel(conn),
	}
}
