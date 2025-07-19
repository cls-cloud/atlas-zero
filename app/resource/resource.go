package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest/httpx"
	"toolkit/helper"
	middle "toolkit/pkg/middleware"
	"toolkit/utils"

	"resource/internal/config"
	"resource/internal/handler"
	"resource/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/resource.yaml", "the config file")

func main() {
	flag.Parse()
	if err := utils.Init("2024-01-01", 1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 创建服务器并传入自定义的 UnauthorizedCallback
	server := rest.MustNewServer(c.RestConf)

	// 使用拦截器
	httpx.SetOkHandler(helper.OkHandler)
	httpx.SetErrorHandlerCtx(helper.ErrHandler(c.RestConf.Name))

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	// 注册中间件
	server.Use(middle.CorsMiddleware)

	group := service.NewServiceGroup()
	group.Add(server)
	defer group.Stop()
	fmt.Printf("Starting server at %s:%d...\n", c.RestConf.Host, c.RestConf.Port)
	group.Start()
}
