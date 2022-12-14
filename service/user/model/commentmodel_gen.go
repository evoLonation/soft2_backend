// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	commentFieldNames          = builder.RawFieldNames(&Comment{})
	commentRows                = strings.Join(commentFieldNames, ",")
	commentRowsExpectAutoSet   = strings.Join(stringx.Remove(commentFieldNames, "`comment_id`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), ",")
	commentRowsWithPlaceHolder = strings.Join(stringx.Remove(commentFieldNames, "`comment_id`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), "=?,") + "=?"
)

type (
	commentModel interface {
		Insert(ctx context.Context, data *Comment) (sql.Result, error)
		FindOne(ctx context.Context, commentId int64) (*Comment, error)
		FindByPaperId(ctx context.Context, paperId string) ([]Comment, error)
		Update(ctx context.Context, data *Comment) error
		Delete(ctx context.Context, commentId int64) error
	}

	defaultCommentModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Comment struct {
		CommentId    int64     `db:"comment_id"`
		UserId       int64     `db:"user_id"`
		UserNickname string    `db:"user_nickname"`
		PaperId      string    `db:"paper_id"`
		Content      string    `db:"content"`
		Likes        int64     `db:"likes"`
		CreateTime   time.Time `db:"create_time"`
	}
)

func newCommentModel(conn sqlx.SqlConn) *defaultCommentModel {
	return &defaultCommentModel{
		conn:  conn,
		table: "`comment`",
	}
}

func (m *defaultCommentModel) Delete(ctx context.Context, commentId int64) error {
	query := fmt.Sprintf("delete from %s where `comment_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, commentId)
	return err
}

func (m *defaultCommentModel) FindOne(ctx context.Context, commentId int64) (*Comment, error) {
	query := fmt.Sprintf("select %s from %s where `comment_id` = ? limit 1", commentRows, m.table)
	var resp Comment
	err := m.conn.QueryRowCtx(ctx, &resp, query, commentId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCommentModel) FindByPaperId(ctx context.Context, paperId string) ([]Comment, error) {
	var resp []Comment
	var query string
	//query = fmt.Sprintf("select %s from %s where `paper_id` = ? limit 1", commentRows, m.table)
	query = fmt.Sprintf("select %s from %s where paper_id = %s", commentRows, m.table, paperId)
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

func (m *defaultCommentModel) Insert(ctx context.Context, data *Comment) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, commentRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.UserNickname, data.PaperId, data.Content, data.Likes)
	return ret, err
}

func (m *defaultCommentModel) Update(ctx context.Context, data *Comment) error {
	query := fmt.Sprintf("update %s set %s where `comment_id` = ?", m.table, commentRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserId, data.UserNickname, data.PaperId, data.Content, data.Likes, data.CommentId)
	return err
}

func (m *defaultCommentModel) tableName() string {
	return m.table
}
