package main

import (
	"context"
	"flag"
	"os"

	"go-mindoc/app/mqueue/cmd/job/internal/config"
	"go-mindoc/app/mqueue/cmd/job/internal/logic"
	"go-mindoc/app/mqueue/cmd/job/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/mqueue.yaml", "Specify the config file")

func main() {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFile, &c)

	// log、prometheus、trace、metricsUrl
	if err := c.SetUp(); err != nil {
		panic(err)
	}

	// logx.DisableStat()

	// load config file to svc.ServiceContext and asynq.Server
	// yaml配置文件内容 解析到 svc.ServiceContext 并 初使化asynq.Server链接
	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()

	// newCronJob class init ctx and all config
	cronJob := logic.NewCronJob(ctx, svcContext)

	// register job  consumer Handle
	mux := cronJob.Register()

	// run AsynqServer consumer
	if err := svcContext.AsynqServer.Run(mux); err != nil {
		logx.WithContext(ctx).Errorf("!!!CronJobErr!!! run err:%+v", err)
		os.Exit(1)
	}
}
