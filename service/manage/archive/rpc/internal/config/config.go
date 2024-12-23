package config

import (
	"Scheduler_go/common/tdenginex"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	Tdengine tdenginex.TDengineConfig
}
