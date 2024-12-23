package svc

import (
	"Scheduler_go/common/tdenginex"
	"Scheduler_go/service/manage/archive/rpc/internal/config"
	"database/sql"
)

type ServiceContext struct {
	Config config.Config

	// Td连接
	Taos *sql.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Taos:   tdenginex.NewTDengineManager(c.Tdengine),
	}
}
