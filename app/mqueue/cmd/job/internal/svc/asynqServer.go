package svc

import (
	"log"

	"github.com/haimait/go-mindoc/app/mqueue/cmd/job/internal/config"

	"github.com/hibiken/asynq"
)

func newAsynqServer(c config.Config) *asynq.Server {

	return asynq.NewServer(
		asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass},
		asynq.Config{
			Concurrency: 20, //max concurrent process job task num 每个进程并发执行的worker数量
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			IsFailure: func(err error) bool {
				log.Printf("asynq server exec task IsFailure ======== >>>>>>>>>>>  err : %+v \n", err)
				return true
			},
		},
	)
}
