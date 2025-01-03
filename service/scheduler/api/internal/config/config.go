package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	ArchiveRpc        zrpc.RpcClientConf
	SchedulerRpc      zrpc.RpcClientConf
	AuthenticationRpc zrpc.RpcClientConf
}
