// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	commentFieldNames          = builder.RawFieldNames(&Comment{})
	commentRows                = strings.Join(commentFieldNames, ",")
	commentRowsExpectAutoSet   = strings.Join(stringx.Remove(commentFieldNames, "`comment_id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), ",")
	commentRowsWithPlaceHolder = strings.Join(stringx.Remove(commentFieldNames, "`comment_id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), "=?,") + "=?"

	cacheCommentCommentIdPrefix = "cache:comment:commentId:"
)

type (
	commentModel interface {
		Insert(ctx context.Context, data *Comment) (sql.Result, error)
		FindOne(ctx context.Context, commentId int64) (*Comment, error)
		FindCommentId(ctx context.Context, userId int64, paperId int64) (*Collect, error)
		Update(ctx context.Context, data *Comment) error
		Delete(ctx context.Context, commentId int64) error
	}

	defaultCommentModel struct {
		sqlc.CachedConn
		table string
	}

	Comment struct {
		CommentId  int64     `db:"comment_id"`
		UserId     int64     `db:"user_id"`
		PaperId    int64     `db:"paper_id"`
		Content    string    `db:"content"`
		Likes      int64     `db:"likes"`
		CreateTime time.Time `db:"create_time"`
	}
)

func newCommentModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultCommentModel {
	return &defaultCommentModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`comment`",
	}
}

func (m *defaultCommentModel) Delete(ctx context.Context, commentId int64) error {
	commentCommentIdKey := fmt.Sprintf("%s%v", cacheCommentCommentIdPrefix, commentId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `comment_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, commentId)
	}, commentCommentIdKey)
	return err
}

func (m *defaultCommentModel) FindOne(ctx context.Context, commentId int64) (*Comment, error) {
	commentCommentIdKey := fmt.Sprintf("%s%v", cacheCommentCommentIdPrefix, commentId)
	var resp Comment
	err := m.QueryRowCtx(ctx, &resp, commentCommentIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `comment_id` = ? limit 1", commentRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, commentId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultCollectModel) FindCommentId(ctx context.Context, userId int64, paperId int64) (*Collect, error) {
	collectCollectIdKey := fmt.Sprintf("%s%v%v", cacheCollectCollectIdPrefix, userId, paperId)
	var resp Collect
	err := m.QueryRowCtx(ctx, &resp, collectCollectIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? and `paper_id` =? limit 2", collectRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, userId, paperId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultCommentModel) Insert(ctx context.Context, data *Comment) (sql.Result, error) {
	commentCommentIdKey := fmt.Sprintf("%s%v", cacheCommentCommentIdPrefix, data.CommentId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, commentRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.PaperId, data.Content, data.Likes)
	}, commentCommentIdKey)
	return ret, err
}

func (m *defaultCommentModel) Update(ctx context.Context, data *Comment) error {
	commentCommentIdKey := fmt.Sprintf("%s%v", cacheCommentCommentIdPrefix, data.CommentId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `comment_id` = ?", m.table, commentRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.PaperId, data.Content, data.Likes, data.CommentId)
	}, commentCommentIdKey)
	return err
}

func (m *defaultCommentModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheCommentCommentIdPrefix, primary)
}

func (m *defaultCommentModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `comment_id` = ? limit 1", commentRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultCommentModel) tableName() string {
	return m.table
}
