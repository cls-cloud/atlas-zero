package main

import (
	"flag"
	"fmt"
	"toolkit/helper"
	"toolkit/utils"

	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest/httpx"

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
	server := rest.MustNewServer(c.RestConf, rest.WithCors("*"))

	// 使用拦截器
	httpx.SetOkHandler(helper.OkHandler)
	httpx.SetErrorHandlerCtx(helper.ErrHandler(c.RestConf.Name))

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	group := service.NewServiceGroup()
	group.Add(server)
	defer group.Stop()
	fmt.Printf("Starting server at %s:%d...\n", c.RestConf.Host, c.RestConf.Port)
	group.Start()
}
