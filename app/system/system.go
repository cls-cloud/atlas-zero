package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest/httpx"
	"system/internal/config"
	"system/internal/handler"
	"system/internal/middleware"
	"system/internal/svc"
	"toolkit/helper"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/system.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	//server := rest.MustNewServer(c.RestConf)
	// 创建服务器并传入自定义的 UnauthorizedCallback
	server := rest.MustNewServer(c.RestConf)
	server.Use(middleware.LogMiddleware)
	server.Use(middleware.ApiMiddleware(c))
	server.Use(middleware.CorsMiddleware)
	// 使用拦截器
	httpx.SetOkHandler(helper.OkHandler)
	httpx.SetErrorHandlerCtx(helper.ErrHandler(c.RestConf.Name))

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	//rpc := zrpc.MustNewServer(c.RpcConf, func(grpcServer *grpc.Server) {
	//	system.RegisterUserRpcServer(grpcServer, userServer.NewUserRpcServer(ctx))
	//	system.RegisterRoleRpcServer(grpcServer, roleServer.NewRoleRpcServer(ctx))
	//
	//	if c.RpcConf.Mode == service.DevMode || c.RpcConf.Mode == service.TestMode {
	//		reflection.Register(grpcServer)
	//	}
	//})
	group := service.NewServiceGroup()
	group.Add(server)
	//group.Add(rpc)
	defer group.Stop()
	fmt.Printf("Starting server at %s:%d...\n", c.RestConf.Host, c.RestConf.Port)
	//fmt.Printf("Starting rpc server at %s...\n", c.RpcConf.ListenOn)
	group.Start()
}
