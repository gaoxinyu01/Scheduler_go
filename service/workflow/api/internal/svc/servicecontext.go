package svc

import (
	"Scheduler_go/service/manage/archive/rpc/archive"
	"Scheduler_go/service/manage/authentication/authentication"
	"Scheduler_go/service/workflow/api/internal/config"
	"Scheduler_go/service/workflow/rpc/workflow"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	WorkflowRpc       workflow.Workflow
	ArchiveRpc        archive.Archive
	AuthenticationRpc authentication.Authentication
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		WorkflowRpc:       workflow.NewWorkflow(zrpc.MustNewClient(c.WorkflowRpc)),
		ArchiveRpc:        archive.NewArchive(zrpc.MustNewClient(c.ArchiveRpc)),
		AuthenticationRpc: authentication.NewAuthentication(zrpc.MustNewClient(c.AuthenticationRpc)),
	}
}
