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
	userHelpFieldNames          = builder.RawFieldNames(&UserHelp{})
	userHelpRows                = strings.Join(userHelpFieldNames, ",")
	userHelpRowsExpectAutoSet   = strings.Join(stringx.Remove(userHelpFieldNames, "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), ",")
	userHelpRowsWithPlaceHolder = strings.Join(stringx.Remove(userHelpFieldNames, "`user_id`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), "=?,") + "=?"
)

type (
	userHelpModel interface {
		Insert(ctx context.Context, data *UserHelp) (sql.Result, error)
		FindOne(ctx context.Context, userId int64) (*UserHelp, error)
		Update(ctx context.Context, data *UserHelp) error
		Delete(ctx context.Context, userId int64) error
	}

	defaultUserHelpModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UserHelp struct {
		UserId  int64 `db:"user_id"` // 用户id
		Request int64 `db:"request"` // 求助次数
		Help    int64 `db:"help"`    // 应助次数
		Wealth  int64 `db:"wealth"`  // 财富值
	}
)

func newUserHelpModel(conn sqlx.SqlConn) *defaultUserHelpModel {
	return &defaultUserHelpModel{
		conn:  conn,
		table: "`user_help`",
	}
}

func (m *defaultUserHelpModel) Delete(ctx context.Context, userId int64) error {
	query := fmt.Sprintf("delete from %s where `user_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, userId)
	return err
}

func (m *defaultUserHelpModel) FindOne(ctx context.Context, userId int64) (*UserHelp, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userHelpRows, m.table)
	var resp UserHelp
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

func (m *defaultUserHelpModel) Insert(ctx context.Context, data *UserHelp) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, userHelpRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.Request, data.Help, data.Wealth)
	return ret, err
}

func (m *defaultUserHelpModel) Update(ctx context.Context, data *UserHelp) error {
	query := fmt.Sprintf("update %s set %s where `user_id` = ?", m.table, userHelpRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Request, data.Help, data.Wealth, data.UserId)
	return err
}

func (m *defaultUserHelpModel) tableName() string {
	return m.table
}
