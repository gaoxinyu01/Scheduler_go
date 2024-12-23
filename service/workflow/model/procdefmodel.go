package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProcDefModel = (*customProcDefModel)(nil)

type (
	// ProcDefModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProcDefModel.

	ProcDefModel interface {
		procDefModel
	}

	customProcDefModel struct {
		*defaultProcDefModel
	}
)

// NewProcDefModel returns a modelx for the database table.
func NewProcDefModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProcDefModel {
	return &customProcDefModel{
		defaultProcDefModel: newProcDefModel(conn, c, opts...),
	}
}
