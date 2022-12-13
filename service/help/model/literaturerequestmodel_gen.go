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
	literatureRequestFieldNames          = builder.RawFieldNames(&LiteratureRequest{})
	literatureRequestRows                = strings.Join(literatureRequestFieldNames, ",")
	literatureRequestRowsExpectAutoSet   = strings.Join(stringx.Remove(literatureRequestFieldNames, "`id`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), ",")
	literatureRequestRowsWithPlaceHolder = strings.Join(stringx.Remove(literatureRequestFieldNames, "`id`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), "=?,") + "=?"
)

type (
	literatureRequestModel interface {
		Insert(ctx context.Context, data *LiteratureRequest) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*LiteratureRequest, error)
		FindAll(ctx context.Context, order int64) ([]LiteratureRequest, error)
		FindByContent(ctx context.Context, content string) ([]LiteratureRequest, error)
		FindComplaint(ctx context.Context) ([]LiteratureRequest, error)
		FindByUser(ctx context.Context, userId int64, status int64) ([]LiteratureRequest, error)
		Update(ctx context.Context, data *LiteratureRequest) error
		Delete(ctx context.Context, id int64) error
	}

	defaultLiteratureRequestModel struct {
		conn  sqlx.SqlConn
		table string
	}

	LiteratureRequest struct {
		Id             int64     `db:"id"`
		UserId         int64     `db:"user_id"`         // 求助者id
		Title          string    `db:"title"`           // 文献标题
		Author         string    `db:"author"`          // 作者名称
		Magazine       string    `db:"magazine"`        // 期刊名称
		Link           string    `db:"link"`            // 链接
		RequestTime    time.Time `db:"request_time"`    // 求助时间
		RequestContent string    `db:"request_content"` // 求助描述
		Wealth         int64     `db:"wealth"`          // 财富值
		RequestStatus  int64     `db:"request_status"`  // 求助状态
		Complaint      string    `db:"complaint"`       // 投诉内容
	}
)

func newLiteratureRequestModel(conn sqlx.SqlConn) *defaultLiteratureRequestModel {
	return &defaultLiteratureRequestModel{
		conn:  conn,
		table: "`literature_request`",
	}
}

func (m *defaultLiteratureRequestModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultLiteratureRequestModel) FindOne(ctx context.Context, id int64) (*LiteratureRequest, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", literatureRequestRows, m.table)
	var resp LiteratureRequest
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

func (m *defaultLiteratureRequestModel) FindAll(ctx context.Context, order int64) ([]LiteratureRequest, error) {
	var resp []LiteratureRequest
	var query string
	if order == 0 {
		query = fmt.Sprintf("select %s from %s where request_status = 0 order by request_time desc", literatureRequestRows, m.table)
	} else {
		query = fmt.Sprintf("select %s from %s where request_status = 0 order by wealth desc", literatureRequestRows, m.table)
	}
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

func (m *defaultLiteratureRequestModel) FindByContent(ctx context.Context, content string) ([]LiteratureRequest, error) {
	var resp []LiteratureRequest
	var query string
	query = fmt.Sprintf("(select %s from %s where title = %s and request_status = 0) union (select %s from %s where request_content = %s and request_status = 0)", literatureRequestRows, m.table, content, literatureRequestRows, m.table, content)
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

func (m *defaultLiteratureRequestModel) FindComplaint(ctx context.Context) ([]LiteratureRequest, error) {
	var resp []LiteratureRequest
	var query string
	query = fmt.Sprintf("select %s from %s where request_status = 3", literatureRequestRows, m.table)
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

func (m *defaultLiteratureRequestModel) FindByUser(ctx context.Context, userId int64, status int64) ([]LiteratureRequest, error) {
	var resp []LiteratureRequest
	var query string
	if status == 0 {
		query = fmt.Sprintf("select %s from %s where user_id = %d", literatureRequestRows, m.table, userId)
	} else {
		query = fmt.Sprintf("select %s from %s where user_id = %d and request_status = %d", literatureRequestRows, m.table, userId, status-1)
	}
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

func (m *defaultLiteratureRequestModel) Insert(ctx context.Context, data *LiteratureRequest) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, literatureRequestRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserId, data.Title, data.Author, data.Magazine, data.Link, data.RequestTime, data.RequestContent, data.Wealth, data.RequestStatus, data.Complaint)
	return ret, err
}

func (m *defaultLiteratureRequestModel) Update(ctx context.Context, data *LiteratureRequest) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, literatureRequestRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.UserId, data.Title, data.Author, data.Magazine, data.Link, data.RequestTime, data.RequestContent, data.Wealth, data.RequestStatus, data.Complaint, data.Id)
	return err
}

func (m *defaultLiteratureRequestModel) tableName() string {
	return m.table
}
