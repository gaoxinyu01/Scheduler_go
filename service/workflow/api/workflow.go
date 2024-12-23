package main

import (
	"flag"
	"fmt"

	"Scheduler_go/service/workflow/api/internal/config"
	"Scheduler_go/service/workflow/api/internal/handler"
	"Scheduler_go/service/workflow/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	zero_handler "github.com/zeromicro/go-zero/rest/handler"
)

var configFile = flag.String("f", "etc/workflow-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf,
		// token错误拦截
		rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
			httpx.WriteJson(w, http.StatusOK, common.NewCodeError(common.TokenErrorCode, msg.TokenError, err.Error()))
		}),
		// 请求方式错误拦截
		rest.WithNotAllowedHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			httpx.WriteJson(w, http.StatusOK, common.NewCodeError(common.ReqNotAllCode, msg.ReqNotAllError, nil))
		})),
		// 路由错误拦截
		rest.WithNotFoundHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			httpx.WriteJson(w, http.StatusOK, common.NewCodeError(common.ReqRoutesErrorCode, msg.ReqRoutesError, nil))
		})),
	)

	defer server.Stop()

	// 设置日志输出 接口慢时间
	zrpc.SetClientSlowThreshold(time.Second * 2000)
	zero_handler.SetSlowThreshold(time.Second * 2000)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
