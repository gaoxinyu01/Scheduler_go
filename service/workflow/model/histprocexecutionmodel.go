package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HistProcExecutionModel = (*customHistProcExecutionModel)(nil)

type (
	// HistProcExecutionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHistProcExecutionModel.

	HistProcExecutionModel interface {
		histProcExecutionModel
	}

	customHistProcExecutionModel struct {
		*defaultHistProcExecutionModel
	}
)

// NewHistProcExecutionModel returns a modelx for the database table.
func NewHistProcExecutionModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) HistProcExecutionModel {
	return &customHistProcExecutionModel{
		defaultHistProcExecutionModel: newHistProcExecutionModel(conn, c, opts...),
	}
}
