package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SchedulingModel = (*customSchedulingModel)(nil)

type (
	// SchedulingModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSchedulingModel.

	SchedulingModel interface {
		schedulingModel
	}

	customSchedulingModel struct {
		*defaultSchedulingModel
	}
)

// NewSchedulingModel returns a modelx for the database table.
func NewSchedulingModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SchedulingModel {
	return &customSchedulingModel{
		defaultSchedulingModel: newSchedulingModel(conn, c, opts...),
	}
}
