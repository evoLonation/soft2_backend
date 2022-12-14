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
	grievanceFieldNames          = builder.RawFieldNames(&Grievance{})
	grievanceRows                = strings.Join(grievanceFieldNames, ",")
	grievanceRowsExpectAutoSet   = strings.Join(stringx.Remove(grievanceFieldNames, "`grievance_id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), ",")
	grievanceRowsWithPlaceHolder = strings.Join(stringx.Remove(grievanceFieldNames, "`grievance_id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), "=?,") + "=?"
)

type (
	grievanceModel interface {
		Insert(ctx context.Context, data *Grievance) (sql.Result, error)
		FindOne(ctx context.Context, grievanceId int64) (*Grievance, error)
		Update(ctx context.Context, data *Grievance) error
		Delete(ctx context.Context, grievanceId int64) error
	}

	defaultGrievanceModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Grievance struct {
		GrievanceId int64  `db:"grievance_id"`
		PlaintiffId string `db:"plaintiff_id"` // 申诉学者id
		DefendantId string `db:"defendant_id"` // 被申诉学者id
		PaperId     string `db:"paper_id"`
	}
)

func newGrievanceModel(conn sqlx.SqlConn) *defaultGrievanceModel {
	return &defaultGrievanceModel{
		conn:  conn,
		table: "`grievance`",
	}
}

func (m *defaultGrievanceModel) Delete(ctx context.Context, grievanceId int64) error {
	query := fmt.Sprintf("delete from %s where `grievance_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, grievanceId)
	return err
}

func (m *defaultGrievanceModel) FindOne(ctx context.Context, grievanceId int64) (*Grievance, error) {
	query := fmt.Sprintf("select %s from %s where `grievance_id` = ? limit 1", grievanceRows, m.table)
	var resp Grievance
	err := m.conn.QueryRowCtx(ctx, &resp, query, grievanceId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultGrievanceModel) Insert(ctx context.Context, data *Grievance) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, grievanceRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.PlaintiffId, data.DefendantId, data.PaperId)
	return ret, err
}

func (m *defaultGrievanceModel) Update(ctx context.Context, data *Grievance) error {
	query := fmt.Sprintf("update %s set %s where `grievance_id` = ?", m.table, grievanceRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.PlaintiffId, data.DefendantId, data.PaperId, data.GrievanceId)
	return err
}

func (m *defaultGrievanceModel) tableName() string {
	return m.table
}
