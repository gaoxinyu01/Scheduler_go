package svc

import (
	"Scheduler_go/service/manage/archive/rpc/archive"
	"Scheduler_go/service/manage/authentication/authentication"
	"Scheduler_go/service/scheduler/api/internal/config"
	"Scheduler_go/service/scheduler/api/internal/middleware"
	"Scheduler_go/service/scheduler/rpc/scheduler"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config            config.Config
	SchedulerRpc      scheduler.Scheduler
	CheckAuth         rest.Middleware
	ArchiveRpc        archive.Archive
	AuthenticationRpc authentication.Authentication
}

func NewServiceContext(c config.Config) *ServiceContext {
	c.ArchiveRpc.Timeout = 30000
	return &ServiceContext{
		Config:            c,
		SchedulerRpc:      scheduler.NewScheduler(zrpc.MustNewClient(c.SchedulerRpc)),
		ArchiveRpc:        archive.NewArchive(zrpc.MustNewClient(c.ArchiveRpc)),
		AuthenticationRpc: authentication.NewAuthentication(zrpc.MustNewClient(c.AuthenticationRpc)),
		CheckAuth:         middleware.NewCheckAuthMiddleware().Handle,
	}
}
