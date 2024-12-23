// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.2

package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	cacheSysUserAccountPrefix = "cache:sysUser:account:"
)

type (
	sysUserModelExtra interface {
		FindByAccount(ctx context.Context, account string) (*SysUser, error)
	}
)

func (m *defaultSysUserModel) FindByAccount(ctx context.Context, account string) (*SysUser, error) {
	sysUserAccountKey := fmt.Sprintf("%s%v", cacheSysUserAccountPrefix, account)
	var resp SysUser
	err := m.QueryRowCtx(ctx, &resp, sysUserAccountKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where deleted_at is null And  `account` = ? limit 1", sysUserRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, account)
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
