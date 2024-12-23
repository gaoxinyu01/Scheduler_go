package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TeamModel = (*customTeamModel)(nil)

type (
	// TeamModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTeamModel.

	TeamModel interface {
		teamModel
	}

	customTeamModel struct {
		*defaultTeamModel
	}
)

// NewTeamModel returns a modelx for the database table.
func NewTeamModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TeamModel {
	return &customTeamModel{
		defaultTeamModel: newTeamModel(conn, c, opts...),
	}
}
