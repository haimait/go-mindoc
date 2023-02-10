package main

import (
	"flag"
	"fmt"
	"go-mindoc/app/api/internal/config"
	"go-mindoc/app/api/internal/handler"
	"go-mindoc/app/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/core-api.yaml", "the config file")

func main() {
	flag.Parse()

	// load config file
	var c config.Config
	conf.MustLoad(*configFile, &c)

	// new newServer
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// load config file to svc.ServiceContext
	ctx := svc.NewServiceContext(c)

	// add routes and handles
	handler.RegisterHandlers(server, ctx)

	// add static path
	handler.SaticRegisterHandlers(server, ctx)

	// print routes
	server.PrintRoutes()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
