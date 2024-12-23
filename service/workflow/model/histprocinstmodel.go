package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ HistProcInstModel = (*customHistProcInstModel)(nil)

type (
	// HistProcInstModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHistProcInstModel.

	HistProcInstModel interface {
		histProcInstModel
	}

	customHistProcInstModel struct {
		*defaultHistProcInstModel
	}
)

// NewHistProcInstModel returns a model for the database table.
func NewHistProcInstModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) HistProcInstModel {
	return &customHistProcInstModel{
		defaultHistProcInstModel: newHistProcInstModel(conn, c, opts...),
	}
}
