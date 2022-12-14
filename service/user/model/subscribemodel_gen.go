// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	subscribeFieldNames          = builder.RawFieldNames(&Subscribe{})
	subscribeRows                = strings.Join(subscribeFieldNames, ",")
	subscribeRowsExpectAutoSet   = strings.Join(stringx.Remove(subscribeFieldNames, "`subscribe_id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), ",")
	subscribeRowsWithPlaceHolder = strings.Join(stringx.Remove(subscribeFieldNames, "`subscribe_id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), "=?,") + "=?"
)

type (
	subscribeModel interface {
		Insert(ctx context.Context, data *Subscribe) (sql.Result, error)
		FindOne(ctx context.Context, subscribeId int64) (*Subscribe, error)
		FindSubscribeId(ctx context.Context, userId int64, scholarId string) (*Subscribe, error)
		FindByUserId(ctx context.Context, userId int64) ([]Subscribe, error)
		Update(ctx context.Context, data *Subscribe) error
		Delete(ctx context.Context, subscribeId int64) error
	}

	defaultSubscribeModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Subscribe struct {
		SubscribeId int64  `db:"subscribe_id"`
		UserId      int64  `db:"user_id"`
		ScholarId   string `db:"scholar_id"`
	}
)

func newSubscribeModel(conn sqlx.SqlConn) *defaultSubscribeModel {
	return &defaultSubscribeModel{
		conn:  conn,
		table: "`subscribe`",
	}
}

func (m *defaultSubscribeModel) Delete(ctx context.Context, subscribeId int64) error {
	query := fmt.Sprintf("delete from %s where `subscribe_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, subscribeId)
	return err
}

func (m *defaultSubscribeModel) FindOne(ctx context.Context, subscribeId int64) (*Subscribe, error) {
	query := fmt.Sprintf("select %s from %s where `subscribe_id` = ? limit 1", subscribeRows, m.table)
	var resp Subscribe
	err := m.conn.QueryRowCtx(ctx, &resp, query, subscribeId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSubscribeModel) FindSubscribeId(ctx context.Context, userId int64, scholarId string) (*Subscribe, error) {
	//query := fmt.Sprintf("select %s from %s where user_id = %d and scholar_id =%s", subscribeRows, m.table, userId, scholarId)
	query := fmt.Sprintf("select %s from %s where `user_id` = ? and `scholar_id` = ? limit 2", subscribeRows, m.table)
	var resp Subscribe
	err := m.conn.QueryRowCtx(ctx, &resp, query, userId, scholarId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSubscribeModel) FindByUserId(ctx context.Context, userId int64) ([]Subscribe, error) {
	var resp []Subscribe
	query := fmt.Sprintf("select %s from %s where user_id = %d", subscribeRows, m.table, userId)
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

func (m *defaultSubscribeModel) Insert(ctx context.Context, data *Subscribe) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, subscribeRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.ScholarId)
	return ret, err
}

func (m *defaultSubscribeModel) Update(ctx context.Context, data *Subscribe) error {
	query := fmt.Sprintf("update %s set %s where `subscribe_id` = ?", m.table, subscribeRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserId, data.ScholarId, data.SubscribeId)
	return err
}

func (m *defaultSubscribeModel) tableName() string {
	return m.table
}
