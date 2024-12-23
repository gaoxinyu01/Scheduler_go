package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HistProcInstVariableModel = (*customHistProcInstVariableModel)(nil)

type (
	// HistProcInstVariableModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHistProcInstVariableModel.

	HistProcInstVariableModel interface {
		histProcInstVariableModel
	}

	customHistProcInstVariableModel struct {
		*defaultHistProcInstVariableModel
	}
)

// NewHistProcInstVariableModel returns a model for the database table.
func NewHistProcInstVariableModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) HistProcInstVariableModel {
	return &customHistProcInstVariableModel{
		defaultHistProcInstVariableModel: newHistProcInstVariableModel(conn, c, opts...),
	}
}
