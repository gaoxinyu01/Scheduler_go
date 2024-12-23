package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HistProcTaskModel = (*customHistProcTaskModel)(nil)

type (
	// HistProcTaskModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHistProcTaskModel.

	HistProcTaskModel interface {
		histProcTaskModel
	}

	customHistProcTaskModel struct {
		*defaultHistProcTaskModel
	}
)

// NewHistProcTaskModel returns a model for the database table.
func NewHistProcTaskModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) HistProcTaskModel {
	return &customHistProcTaskModel{
		defaultHistProcTaskModel: newHistProcTaskModel(conn, c, opts...),
	}
}
