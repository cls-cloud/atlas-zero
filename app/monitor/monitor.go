package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	logininforpcServer "monitor/internal/server/logininforpc"
	operlogrpcServer "monitor/internal/server/operlogrpc"
	"monitor/pb/monitor"
	"toolkit/helper"
	"toolkit/middlewares"
	middle "toolkit/pkg/middleware"
	"toolkit/utils"

	"monitor/internal/config"
	"monitor/internal/handler"
	"monitor/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/monitor.yaml", "the config file")

func main() {
	flag.Parse()
	if err := utils.Init("2024-01-01", 1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)

	httpx.SetOkHandler(helper.OkHandler)
	httpx.SetErrorHandlerCtx(helper.ErrHandler(c.RestConf.Name))

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	// 注册中间件
	//server.Use(middleware.LogMiddleware)
	server.Use(middlewares.ApiMiddleware(c.RestConf.Mode))
	server.Use(middle.CorsMiddleware)

	rpc := zrpc.MustNewServer(c.RpcConf, func(grpcServer *grpc.Server) {
		monitor.RegisterLoginInfoRpcServer(grpcServer, logininforpcServer.NewLoginInfoRpcServer(ctx))
		monitor.RegisterOperLogRpcServer(grpcServer, operlogrpcServer.NewOperLogRpcServer(ctx))

		if c.RpcConf.Mode == service.DevMode || c.RpcConf.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	group := service.NewServiceGroup()
	group.Add(server)
	group.Add(rpc)
	defer group.Stop()
	fmt.Printf("Starting server at %s:%d...\n", c.RestConf.Host, c.RestConf.Port)
	fmt.Printf("Starting rpc server at %s...\n", c.RpcConf.ListenOn)
	group.Start()
}
