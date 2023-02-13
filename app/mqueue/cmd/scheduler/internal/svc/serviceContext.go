package svc

import (
	"github.com/hibiken/asynq"

	"go-mindoc/app/mqueue/cmd/scheduler/internal/config"
)

type ServiceContext struct {
	Config config.Config

	AsynqScheduler *asynq.Scheduler
	AsynqClient    *asynq.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		AsynqScheduler: newScheduler(c),
		AsynqClient:    newAsynqClient(c),
	}
}
