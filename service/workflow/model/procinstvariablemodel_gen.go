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
	procInstVariableFieldNames          = builder.RawFieldNames(&ProcInstVariable{})
	procInstVariableRows                = strings.Join(procInstVariableFieldNames, ",")
	procInstVariableRowsExpectAutoSet   = strings.Join(stringx.Remove(procInstVariableFieldNames, "`id`", "`create_at`", "`create_time`", "`update_at`", "`update_time`"), ",")
	procInstVariableRowsWithPlaceHolder = strings.Join(stringx.Remove(procInstVariableFieldNames, "`id`", "`create_at`", "`create_time`", "`update_at`", "`update_time`"), "=?,") + "=?"

	cacheProcInstVariableIdPrefix = "cache:procInstVariable:id:"
)

type (
	procInstVariableModel interface {
		Insert(ctx context.Context, data *ProcInstVariable) (sql.Result, error)
		TransInsert(ctx context.Context,  session sqlx.Session,data *ProcInstVariable) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ProcInstVariable, error)
		Update(ctx context.Context, data *ProcInstVariable) error
		TransUpdate(ctx context.Context,  session sqlx.Session,data *ProcInstVariable) error
		Delete(ctx context.Context, id int64) error
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindList(ctx context.Context, rowBuilder squirrel.SelectBuilder, current, pageSize int64) ([]*ProcInstVariable, error)
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		TransCtx(ctx context.Context, fn func(ctx context.Context, sqlx sqlx.Session) error) error
		FindOneByProcInstIdAndKey( procInstId int64,key string) (*ProcInstVariable, error)
	}

	defaultProcInstVariableModel struct {
		sqlc.CachedConn
		table string
	}

	ProcInstVariable struct {
		Id          int64          `db:"id"`           // 身份ID
		ProcInstId  int64          `db:"proc_inst_id"` // 流程实例ID
		Key         string         `db:"key"`          // 变量key
		Value       string         `db:"value"`        // 变量value
		TenantId    string         `db:"tenant_id"`    // 租户ID
		Data        sql.NullString `db:"data"`
		CreatedAt   time.Time      `db:"created_at"`   // 创建时间
		UpdatedAt   sql.NullTime   `db:"updated_at"`   // 更新时间
		DeletedAt   sql.NullTime   `db:"deleted_at"`   // 删除时间
		CreatedName string         `db:"created_name"` // 创建人
		UpdatedName sql.NullString `db:"updated_name"` // 更新人
		DeletedName sql.NullString `db:"deleted_name"` // 删除人
	}
)

func newProcInstVariableModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultProcInstVariableModel {
	return &defaultProcInstVariableModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`proc_inst_variable`",
	}
}

func (m *defaultProcInstVariableModel) Delete(ctx context.Context, id int64) error {
	procInstVariableIdKey := fmt.Sprintf("%s%v", cacheProcInstVariableIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, procInstVariableIdKey)
	return err
}

func (m *defaultProcInstVariableModel) FindOne(ctx context.Context, id int64) (*ProcInstVariable, error) {
	procInstVariableIdKey := fmt.Sprintf("%s%v", cacheProcInstVariableIdPrefix, id)
	var resp ProcInstVariable
	err := m.QueryRowCtx(ctx, &resp, procInstVariableIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", procInstVariableRows, m.table)
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

func (m *defaultProcInstVariableModel) FindOneByProcInstIdAndKey( procInstId int64,key string) (*ProcInstVariable, error) {
	var resp ProcInstVariable
	query := fmt.Sprintf("select %s from %s where proc_inst_id=? AND `key`=? limit 1", procInstVariableRows, m.table )
	err := m.QueryRowNoCache( &resp, query)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, sqlx.ErrNotFound
	default:
		return nil, err
	}

}

func (m *defaultProcInstVariableModel) Insert(ctx context.Context, data *ProcInstVariable) (sql.Result, error) {
	procInstVariableIdKey := fmt.Sprintf("%s%v", cacheProcInstVariableIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, procInstVariableRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ProcInstId, data.Key, data.Value, data.TenantId, data.Data, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.CreatedName, data.UpdatedName, data.DeletedName)
	}, procInstVariableIdKey)
	return ret, err
}

func (m *defaultProcInstVariableModel) TransInsert(ctx context.Context,  session sqlx.Session,data *ProcInstVariable) (sql.Result, error) {
	procInstVariableIdKey := fmt.Sprintf("%s%v", cacheProcInstVariableIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, procInstVariableRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.ProcInstId, data.Key, data.Value, data.TenantId, data.Data, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.CreatedName, data.UpdatedName, data.DeletedName)
	}, procInstVariableIdKey)
	return ret, err
}
func (m *defaultProcInstVariableModel) Update(ctx context.Context, data *ProcInstVariable) error {
	procInstVariableIdKey := fmt.Sprintf("%s%v", cacheProcInstVariableIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, procInstVariableRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.ProcInstId, data.Key, data.Value, data.TenantId, data.Data, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.CreatedName, data.UpdatedName, data.DeletedName, data.Id)
	}, procInstVariableIdKey)
	return err
}

func (m *defaultProcInstVariableModel)TransUpdate(ctx context.Context,  session sqlx.Session,data *ProcInstVariable) error {
	procInstVariableIdKey := fmt.Sprintf("%s%v", cacheProcInstVariableIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, procInstVariableRowsWithPlaceHolder)
		return session.ExecCtx(ctx, query, data.ProcInstId, data.Key, data.Value, data.TenantId, data.Data, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.CreatedName, data.UpdatedName, data.DeletedName, data.Id)
	}, procInstVariableIdKey)
	return err
}

func (m *defaultProcInstVariableModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(procInstVariableRows).From(m.table)
}

func (m *defaultProcInstVariableModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultProcInstVariableModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}

func (m *defaultProcInstVariableModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultProcInstVariableModel) FindList(ctx context.Context, rowBuilder squirrel.SelectBuilder, current, pageSize int64) ([]*ProcInstVariable, error) {

	if current < 1 {
		current = 1
	}
	offset := (current - 1) * pageSize

	query, values, err := rowBuilder.Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*ProcInstVariable
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

func (m *defaultProcInstVariableModel) TransCtx(ctx context.Context, fn func(ctx context.Context, sqlx sqlx.Session) error) error {
	return m.Transact(func(s sqlx.Session) error {
		return fn(ctx, s)
	})
}

func (m *defaultProcInstVariableModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheProcInstVariableIdPrefix, primary)
}

func (m *defaultProcInstVariableModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", procInstVariableRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultProcInstVariableModel) tableName() string {
	return m.table
}
