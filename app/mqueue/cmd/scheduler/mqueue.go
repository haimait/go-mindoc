package main

import (
	"context"
	"flag"
	"os"

	"go-mindoc/app/mqueue/cmd/scheduler/internal/config"
	"go-mindoc/app/mqueue/cmd/scheduler/internal/logic"
	"go-mindoc/app/mqueue/cmd/scheduler/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/mqueue.yaml", "Specify the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// DisableStat disables the stat logs.
	logx.DisableStat()

	// log、prometheus、trace、metricsUrl.
	if err := c.SetUp(); err != nil {
		panic(err)
	}

	// load config file to svc.ServiceContext
	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()

	// init ctx and all config
	mqueueScheduler := logic.NewCronScheduler(ctx, svcContext)

	// register job  producer Handle
	mqueueScheduler.RegisterCluster()
	mqueueScheduler.RegisterScheduler()

	// run producer service
	if err := svcContext.AsynqScheduler.Run(); err != nil {
		logx.Errorf("!!!MqueueSchedulerErr!!!  run err:%+v", err)
		os.Exit(1)
	}

}
