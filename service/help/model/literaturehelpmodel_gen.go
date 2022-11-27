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
	literatureHelpFieldNames          = builder.RawFieldNames(&LiteratureHelp{})
	literatureHelpRows                = strings.Join(literatureHelpFieldNames, ",")
	literatureHelpRowsExpectAutoSet   = strings.Join(stringx.Remove(literatureHelpFieldNames, "`id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), ",")
	literatureHelpRowsWithPlaceHolder = strings.Join(stringx.Remove(literatureHelpFieldNames, "`id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), "=?,") + "=?"
)

type (
	literatureHelpModel interface {
		Insert(ctx context.Context, data *LiteratureHelp) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*LiteratureHelp, error)
		Update(ctx context.Context, data *LiteratureHelp) error
		Delete(ctx context.Context, id int64) error
	}

	defaultLiteratureHelpModel struct {
		conn  sqlx.SqlConn
		table string
	}

	LiteratureHelp struct {
		Id         int64 `db:"id"`
		UserId     int64 `db:"user_id"`     // 应助者id
		RequestId  int64 `db:"request_id"`  // 求助id
		Wealth     int64 `db:"wealth"`      // 财富值
		HelpStatus int64 `db:"help_status"` // 应助状态
	}
)

func newLiteratureHelpModel(conn sqlx.SqlConn) *defaultLiteratureHelpModel {
	return &defaultLiteratureHelpModel{
		conn:  conn,
		table: "`literature_help`",
	}
}

func (m *defaultLiteratureHelpModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultLiteratureHelpModel) FindOne(ctx context.Context, id int64) (*LiteratureHelp, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", literatureHelpRows, m.table)
	var resp LiteratureHelp
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLiteratureHelpModel) Insert(ctx context.Context, data *LiteratureHelp) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, literatureHelpRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.RequestId, data.Wealth, data.HelpStatus)
	return ret, err
}

func (m *defaultLiteratureHelpModel) Update(ctx context.Context, data *LiteratureHelp) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, literatureHelpRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserId, data.RequestId, data.Wealth, data.HelpStatus, data.Id)
	return err
}

func (m *defaultLiteratureHelpModel) tableName() string {
	return m.table
}
