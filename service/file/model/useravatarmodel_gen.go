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
	userAvatarFieldNames          = builder.RawFieldNames(&UserAvatar{})
	userAvatarRows                = strings.Join(userAvatarFieldNames, ",")
	userAvatarRowsExpectAutoSet   = strings.Join(stringx.Remove(userAvatarFieldNames, "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`"), ",")
	userAvatarRowsWithPlaceHolder = strings.Join(stringx.Remove(userAvatarFieldNames, "`user_id`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`"), "=?,") + "=?"
)

type (
	userAvatarModel interface {
		Insert(ctx context.Context, data *UserAvatar) (sql.Result, error)
		FindOne(ctx context.Context, userId string) (*UserAvatar, error)
		Update(ctx context.Context, data *UserAvatar) error
		Delete(ctx context.Context, userId string) error
	}

	defaultUserAvatarModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UserAvatar struct {
		UserId   string `db:"user_id"`
		FileName string `db:"file_name"`
	}
)

func newUserAvatarModel(conn sqlx.SqlConn) *defaultUserAvatarModel {
	return &defaultUserAvatarModel{
		conn:  conn,
		table: "`user_avatar`",
	}
}

func (m *defaultUserAvatarModel) Delete(ctx context.Context, userId string) error {
	query := fmt.Sprintf("delete from %s where `user_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, userId)
	return err
}

func (m *defaultUserAvatarModel) FindOne(ctx context.Context, userId string) (*UserAvatar, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userAvatarRows, m.table)
	var resp UserAvatar
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

func (m *defaultUserAvatarModel) Insert(ctx context.Context, data *UserAvatar) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, userAvatarRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.FileName)
	return ret, err
}

func (m *defaultUserAvatarModel) Update(ctx context.Context, data *UserAvatar) error {
	query := fmt.Sprintf("update %s set %s where `user_id` = ?", m.table, userAvatarRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.FileName, data.UserId)
	return err
}

func (m *defaultUserAvatarModel) tableName() string {
	return m.table
}
