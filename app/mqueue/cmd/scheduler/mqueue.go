package main

import (
	"context"
	"flag"
	"os"

	"github.com/haimait/go-mindoc/app/mqueue/cmd/scheduler/internal/config"
	"github.com/haimait/go-mindoc/app/mqueue/cmd/scheduler/internal/logic"
	"github.com/haimait/go-mindoc/app/mqueue/cmd/scheduler/internal/svc"

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

	mqueueScheduler := logic.NewCronScheduler(ctx, svcContext)
	mqueueScheduler.Register()

	if err := svcContext.Scheduler.Run(); err != nil {
		logx.Errorf("!!!MqueueSchedulerErr!!!  run err:%+v", err)
		os.Exit(1)
	}

}
