package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SchedulingTypeModel = (*customSchedulingTypeModel)(nil)

type (
	// SchedulingTypeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSchedulingTypeModel.

	SchedulingTypeModel interface {
		schedulingTypeModel
	}

	customSchedulingTypeModel struct {
		*defaultSchedulingTypeModel
	}
)

// NewSchedulingTypeModel returns a modelx for the database table.
func NewSchedulingTypeModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SchedulingTypeModel {
	return &customSchedulingTypeModel{
		defaultSchedulingTypeModel: newSchedulingTypeModel(conn, c, opts...),
	}
}
