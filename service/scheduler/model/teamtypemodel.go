package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TeamTypeModel = (*customTeamTypeModel)(nil)

type (
	// TeamTypeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTeamTypeModel.

	TeamTypeModel interface {
		teamTypeModel
	}

	customTeamTypeModel struct {
		*defaultTeamTypeModel
	}
)

// NewTeamTypeModel returns a modelx for the database table.
func NewTeamTypeModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TeamTypeModel {
	return &customTeamTypeModel{
		defaultTeamTypeModel: newTeamTypeModel(conn, c, opts...),
	}
}
