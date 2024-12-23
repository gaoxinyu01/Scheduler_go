package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProcExecutionModel = (*customProcExecutionModel)(nil)

type (
	// ProcExecutionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProcExecutionModel.

	ProcExecutionModel interface {
		procExecutionModel
	}

	customProcExecutionModel struct {
		*defaultProcExecutionModel
	}
)

// NewProcExecutionModel returns a modelx for the database table.
func NewProcExecutionModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProcExecutionModel {
	return &customProcExecutionModel{
		defaultProcExecutionModel: newProcExecutionModel(conn, c, opts...),
	}
}
