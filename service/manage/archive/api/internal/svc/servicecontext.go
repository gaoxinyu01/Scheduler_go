package svc

import (
	"Scheduler_go/service/manage/archive/api/internal/config"
	"Scheduler_go/service/manage/archive/rpc/archive"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	// 日志服务
	ArchiveRpc archive.Archive
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ArchiveRpc: archive.NewArchive(zrpc.MustNewClient(c.ArchiveRpc)),
	}
}
