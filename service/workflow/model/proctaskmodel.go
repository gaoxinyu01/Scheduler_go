package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProcTaskModel = (*customProcTaskModel)(nil)

type (
	// ProcTaskModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProcTaskModel.

	ProcTaskModel interface {
		procTaskModel
	}

	customProcTaskModel struct {
		*defaultProcTaskModel
	}
)

// NewProcTaskModel returns a model for the database table.
func NewProcTaskModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProcTaskModel {
	return &customProcTaskModel{
		defaultProcTaskModel: newProcTaskModel(conn, c, opts...),
	}
}
