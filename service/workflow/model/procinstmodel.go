package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProcInstModel = (*customProcInstModel)(nil)

type (
	// ProcInstModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProcInstModel.

	ProcInstModel interface {
		procInstModel
	}

	customProcInstModel struct {
		*defaultProcInstModel
	}
)

// NewProcInstModel returns a modelx for the database table.
func NewProcInstModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProcInstModel {
	return &customProcInstModel{
		defaultProcInstModel: newProcInstModel(conn, c, opts...),
	}
}
