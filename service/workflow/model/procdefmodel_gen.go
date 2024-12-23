// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.2

package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	procDefFieldNames          = builder.RawFieldNames(&ProcDef{})
	procDefRows                = strings.Join(procDefFieldNames, ",")
	procDefRowsExpectAutoSet   = strings.Join(stringx.Remove(procDefFieldNames, "`id`", "`create_at`", "`create_time`", "`update_at`", "`update_time`"), ",")
	procDefRowsWithPlaceHolder = strings.Join(stringx.Remove(procDefFieldNames, "`id`", "`create_at`", "`create_time`", "`update_at`", "`update_time`"), "=?,") + "=?"

	cacheProcDefIdPrefix = "cache:procDef:id:"
	cacheProcDefNamePrefix = "cache:procDef:name:"
)

type (
	procDefModel interface {
		Insert(ctx context.Context, data *ProcDef) (sql.Result, error)
		TransInsert(ctx context.Context,  session sqlx.Session,data *ProcDef) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ProcDef, error)
		Update(ctx context.Context, data *ProcDef) error
		Delete(ctx context.Context, id int64) error
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindList(ctx context.Context, rowBuilder squirrel.SelectBuilder, current, pageSize int64) ([]*ProcDef, error)
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		TransCtx(ctx context.Context, fn func(ctx context.Context, sqlx sqlx.Session) error) error
		FindOneByProcessName(ctx context.Context, name string) (*ProcDef, error)
	}

	defaultProcDefModel struct {
		sqlc.CachedConn
		table string
	}

	ProcDef struct {
		Id           int64          `db:"id"`             // 流程模板ID
		Name         string         `db:"name"`           // 流程名称
		Version      int64          `db:"version"`        // 版本号
		ProcType     int64          `db:"proc_type"`      // 流程类型
		Resource     string         `db:"resource"`       // 流程定义模板
		CreateUserId sql.NullString `db:"create_user_id"` // 创建者ID
		Source       sql.NullString `db:"source"`         // 来源
		TenantId     string         `db:"tenant_id"`      // 租户ID
		Data         sql.NullString `db:"data"`
		CreatedAt    time.Time      `db:"created_at"`   // 创建时间
		UpdatedAt    sql.NullTime   `db:"updated_at"`   // 更新时间
		DeletedAt    sql.NullTime   `db:"deleted_at"`   // 删除时间
		CreatedName  string         `db:"created_name"` // 创建人
		UpdatedName  sql.NullString `db:"updated_name"` // 更新人
		DeletedName  sql.NullString `db:"deleted_name"` // 删除人
	}
)

func newProcDefModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultProcDefModel {
	return &defaultProcDefModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`proc_def`",
	}
}

func (m *defaultProcDefModel) Delete(ctx context.Context, id int64) error {
	procDefIdKey := fmt.Sprintf("%s%v", cacheProcDefIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, procDefIdKey)
	return err
}

func (m *defaultProcDefModel) FindOne(ctx context.Context, id int64) (*ProcDef, error) {
	procDefIdKey := fmt.Sprintf("%s%v", cacheProcDefIdPrefix, id)
	var resp ProcDef
	err := m.QueryRowCtx(ctx, &resp, procDefIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", procDefRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, sqlx.ErrNotFound
	default:
		return nil, err
	}

}
func (m *defaultProcDefModel) FindOneByProcessName(ctx context.Context, name string) (*ProcDef, error) {
	procDefNameKey := fmt.Sprintf("%s%v", cacheProcDefNamePrefix, name)
	var resp ProcDef
	err := m.QueryRowCtx(ctx, &resp, procDefNameKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", procDefRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, name)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, sqlx.ErrNotFound
	default:
		return nil, err
	}

}

func (m *defaultProcDefModel) Insert(ctx context.Context, data *ProcDef) (sql.Result, error) {
	procDefIdKey := fmt.Sprintf("%s%v", cacheProcDefIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, procDefRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name, data.Version, data.ProcType, data.Resource, data.CreateUserId, data.Source, data.TenantId, data.Data, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.CreatedName, data.UpdatedName, data.DeletedName)
	}, procDefIdKey)
	return ret, err
}

func (m *defaultProcDefModel) TransInsert(ctx context.Context,  session sqlx.Session,data *ProcDef) (sql.Result, error) {
	procDefIdKey := fmt.Sprintf("%s%v", cacheProcDefIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, procDefRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.Name, data.Version, data.ProcType, data.Resource, data.CreateUserId, data.Source, data.TenantId, data.Data, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.CreatedName, data.UpdatedName, data.DeletedName)
	}, procDefIdKey)
	return ret, err
}

func (m *defaultProcDefModel) Update(ctx context.Context, data *ProcDef) error {
	procDefIdKey := fmt.Sprintf("%s%v", cacheProcDefIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, procDefRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Name, data.Version, data.ProcType, data.Resource, data.CreateUserId, data.Source, data.TenantId, data.Data, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.CreatedName, data.UpdatedName, data.DeletedName, data.Id)
	}, procDefIdKey)
	return err
}

func (m *defaultProcDefModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(procDefRows).From(m.table)
}

func (m *defaultProcDefModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultProcDefModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}

func (m *defaultProcDefModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

	query, values, err := countBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return 0, nil
	default:
		return 0, err
	}
}

func (m *defaultProcDefModel) FindList(ctx context.Context, rowBuilder squirrel.SelectBuilder, current, pageSize int64) ([]*ProcDef, error) {

	if current < 1 {
		current = 1
	}
	offset := (current - 1) * pageSize

	query, values, err := rowBuilder.Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*ProcDef
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *defaultProcDefModel) TransCtx(ctx context.Context, fn func(ctx context.Context, sqlx sqlx.Session) error) error {
	return m.Transact(func(s sqlx.Session) error {
		return fn(ctx, s)
	})
}

func (m *defaultProcDefModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheProcDefIdPrefix, primary)
}

func (m *defaultProcDefModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", procDefRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultProcDefModel) tableName() string {
	return m.table
}
