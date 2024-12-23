package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProcInstVariableModel = (*customProcInstVariableModel)(nil)

type (
	// ProcInstVariableModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProcInstVariableModel.

	ProcInstVariableModel interface {
		procInstVariableModel
	}

	customProcInstVariableModel struct {
		*defaultProcInstVariableModel
	}
)

// NewProcInstVariableModel returns a model for the database table.
func NewProcInstVariableModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProcInstVariableModel {
	return &customProcInstVariableModel{
		defaultProcInstVariableModel: newProcInstVariableModel(conn, c, opts...),
	}
}
