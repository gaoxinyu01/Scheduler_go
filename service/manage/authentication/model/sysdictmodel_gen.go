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
	sysDictFieldNames          = builder.RawFieldNames(&SysDict{})
	sysDictRows                = strings.Join(sysDictFieldNames, ",")
	sysDictRowsExpectAutoSet   = strings.Join(stringx.Remove(sysDictFieldNames, "`id`", "`create_at`", "`create_time`", "`update_at`", "`update_time`"), ",")
	sysDictRowsWithPlaceHolder = strings.Join(stringx.Remove(sysDictFieldNames, "`id`", "`create_at`", "`create_time`", "`update_at`", "`update_time`"), "=?,") + "=?"

	cacheSysDictIdPrefix = "cache:sysDict:id:"
)

type (
	sysDictModel interface {
		Insert(ctx context.Context, data *SysDict) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysDict, error)
		Update(ctx context.Context, data *SysDict) error
		Delete(ctx context.Context, id int64) error
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindList(ctx context.Context, rowBuilder squirrel.SelectBuilder, current, pageSize int64) ([]*SysDict, error)
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		TransCtx(ctx context.Context, fn func(ctx context.Context, sqlx sqlx.Session) error) error
	}

	defaultSysDictModel struct {
		sqlc.CachedConn
		table string
	}

	SysDict struct {
		Id          int64          `db:"id"`           // 字典类型ID
		CreatedAt   time.Time      `db:"created_at"`   // 创建时间
		UpdatedAt   sql.NullTime   `db:"updated_at"`   // 更新时间
		DeletedAt   sql.NullTime   `db:"deleted_at"`   // 删除时间
		CreatedName string         `db:"created_name"` // 创建人
		UpdatedName sql.NullString `db:"updated_name"` // 更新人
		DeletedName sql.NullString `db:"deleted_name"` // 删除人
		DictType    string         `db:"dict_type"`    // 字典类型
		DictLabel   string         `db:"dict_label"`   // 字典标签
		DictValue   string         `db:"dict_value"`   // 字典键值
		Sort        int64          `db:"sort"`         // 排序
		Remark      sql.NullString `db:"remark"`       // 备注
		State       int64          `db:"state"`        // 状态
	}
)

func newSysDictModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultSysDictModel {
	return &defaultSysDictModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`sys_dict`",
	}
}

func (m *defaultSysDictModel) Delete(ctx context.Context, id int64) error {
	sysDictIdKey := fmt.Sprintf("%s%v", cacheSysDictIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, sysDictIdKey)
	return err
}

func (m *defaultSysDictModel) FindOne(ctx context.Context, id int64) (*SysDict, error) {
	sysDictIdKey := fmt.Sprintf("%s%v", cacheSysDictIdPrefix, id)
	var resp SysDict
	err := m.QueryRowCtx(ctx, &resp, sysDictIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysDictRows, m.table)
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

func (m *defaultSysDictModel) Insert(ctx context.Context, data *SysDict) (sql.Result, error) {
	sysDictIdKey := fmt.Sprintf("%s%v", cacheSysDictIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysDictRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.CreatedName, data.UpdatedName, data.DeletedName, data.DictType, data.DictLabel, data.DictValue, data.Sort, data.Remark, data.State)
	}, sysDictIdKey)
	return ret, err
}

func (m *defaultSysDictModel) Update(ctx context.Context, data *SysDict) error {
	sysDictIdKey := fmt.Sprintf("%s%v", cacheSysDictIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysDictRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.CreatedName, data.UpdatedName, data.DeletedName, data.DictType, data.DictLabel, data.DictValue, data.Sort, data.Remark, data.State, data.Id)
	}, sysDictIdKey)
	return err
}

func (m *defaultSysDictModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(sysDictRows).From(m.table)
}

func (m *defaultSysDictModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultSysDictModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}

func (m *defaultSysDictModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultSysDictModel) FindList(ctx context.Context, rowBuilder squirrel.SelectBuilder, current, pageSize int64) ([]*SysDict, error) {

	if current < 1 {
		current = 1
	}
	offset := (current - 1) * pageSize

	query, values, err := rowBuilder.Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*SysDict
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

func (m *defaultSysDictModel) TransCtx(ctx context.Context, fn func(ctx context.Context, sqlx sqlx.Session) error) error {
	return m.Transact(func(s sqlx.Session) error {
		return fn(ctx, s)
	})
}

func (m *defaultSysDictModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheSysDictIdPrefix, primary)
}

func (m *defaultSysDictModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysDictRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultSysDictModel) tableName() string {
	return m.table
}