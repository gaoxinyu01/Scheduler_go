package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HistProcDefModel = (*customHistProcDefModel)(nil)

type (
	// HistProcDefModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHistProcDefModel.

	HistProcDefModel interface {
		histProcDefModel
	}

	customHistProcDefModel struct {
		*defaultHistProcDefModel
	}
)

// NewHistProcDefModel returns a modelx for the database table.
func NewHistProcDefModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) HistProcDefModel {
	return &customHistProcDefModel{
		defaultHistProcDefModel: newHistProcDefModel(conn, c, opts...),
	}
}
