package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	WorkflowRpc       zrpc.RpcClientConf
	ArchiveRpc        zrpc.RpcClientConf
	SchedulerRpc      zrpc.RpcClientConf
	AuthenticationRpc zrpc.RpcClientConf
}
