package main

import (
	"Scheduler_go/service/manage/archive/rpc/archiveclient"
	"Scheduler_go/service/manage/archive/rpc/internal/config"
	"Scheduler_go/service/manage/archive/rpc/internal/server"
	"Scheduler_go/service/manage/archive/rpc/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"runtime"
)

var configFile = flag.String("f", "etc/archive.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	cpuNum := runtime.NumCPU() //获得当前设备的cpu核心数
	fmt.Println("archiveRpc任务使用,cpu核心数:", cpuNum)
	runtime.GOMAXPROCS(cpuNum) //设置需要用到的cpu数量

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		archiveclient.RegisterArchiveServer(grpcServer, server.NewArchiveServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	//rpc log
	//s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)

	fmt.Println(c.Name, c.ListenOn)
	s.Start()
}
