package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
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
	}

	customApplyModel struct {
		*defaultApplyModel
	}
)

func (m *customApplyModel) FindAll(ctx context.Context) ([]*Apply, error) {
	var resp []*Apply

	query := fmt.Sprintf("select %s from %s where `status` = 0", applyRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewApplyModel returns a model for the database table.
func NewApplyModel(conn sqlx.SqlConn, c cache.CacheConf) ApplyModel {
	return &customApplyModel{
		defaultApplyModel: newApplyModel(conn, c),
	}
}
