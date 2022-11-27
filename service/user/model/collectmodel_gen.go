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
	collectFieldNames          = builder.RawFieldNames(&Collect{})
	collectRows                = strings.Join(collectFieldNames, ",")
	collectRowsExpectAutoSet   = strings.Join(stringx.Remove(collectFieldNames, "`collect_id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), ",")
	collectRowsWithPlaceHolder = strings.Join(stringx.Remove(collectFieldNames, "`collect_id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), "=?,") + "=?"

	cacheCollectCollectIdPrefix = "cache:collect:collectId:"
)

type (
	collectModel interface {
		Insert(ctx context.Context, data *Collect) (sql.Result, error)
		FindOne(ctx context.Context, collectId int64) (*Collect, error)
		FindOneByTwo(ctx context.Context, userId int64, paperId int64) (*Collect, error)
		Update(ctx context.Context, data *Collect) error
		Delete(ctx context.Context, collectId int64) error
	}

	defaultCollectModel struct {
		sqlc.CachedConn
		table string
	}

	Collect struct {
		CollectId  int64     `db:"collect_id"`
		UserId     int64     `db:"user_id"`
		PaperId    int64     `db:"paper_id"`
		CreateTime time.Time `db:"create_time"`
	}
)

func newCollectModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultCollectModel {
	return &defaultCollectModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`collect`",
	}
}

func (m *defaultCollectModel) Delete(ctx context.Context, collectId int64) error {
	collectCollectIdKey := fmt.Sprintf("%s%v", cacheCollectCollectIdPrefix, collectId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `collect_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, collectId)
	}, collectCollectIdKey)
	return err
}

func (m *defaultCollectModel) FindOne(ctx context.Context, collectId int64) (*Collect, error) {
	collectCollectIdKey := fmt.Sprintf("%s%v", cacheCollectCollectIdPrefix, collectId)
	var resp Collect
	err := m.QueryRowCtx(ctx, &resp, collectCollectIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `collect_id` = ? limit 1", collectRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, collectId)
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
func (m *defaultCollectModel) FindOneByTwo(ctx context.Context, userId int64, paperId int64) (*Collect, error) {
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
func (m *defaultCollectModel) Insert(ctx context.Context, data *Collect) (sql.Result, error) {
	collectCollectIdKey := fmt.Sprintf("%s%v", cacheCollectCollectIdPrefix, data.CollectId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, collectRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.PaperId)
	}, collectCollectIdKey)
	return ret, err
}

func (m *defaultCollectModel) Update(ctx context.Context, data *Collect) error {
	collectCollectIdKey := fmt.Sprintf("%s%v", cacheCollectCollectIdPrefix, data.CollectId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `collect_id` = ?", m.table, collectRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.PaperId, data.CollectId)
	}, collectCollectIdKey)
	return err
}

func (m *defaultCollectModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheCollectCollectIdPrefix, primary)
}

func (m *defaultCollectModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `collect_id` = ? limit 1", collectRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultCollectModel) tableName() string {
	return m.table
}