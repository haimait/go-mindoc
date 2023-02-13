package svc

import (
	"fmt"
	"time"

	"go-mindoc/app/mqueue/cmd/scheduler/internal/config"

	"github.com/hibiken/asynq"
)

// create scheduler 初使化创建定时任务 消费者
func newScheduler(c config.Config) *asynq.Scheduler {

	location, _ := time.LoadLocation("Asia/Shanghai")
	return asynq.NewScheduler(
		asynq.RedisClientOpt{
			Addr:     c.Redis.Host,
			Password: c.Redis.Pass,
		}, &asynq.SchedulerOpts{
			Location: location,
			PostEnqueueFunc: func(task *asynq.TaskInfo, err error) {
				if err != nil {
					fmt.Printf("AsynqScheduler EnqueueErrorHandler <<<<<<<===>>>>> err : %+v , task : %+v \n", err, task)
				} else {
					fmt.Printf("AsynqScheduler product <<<<<<<===>>>>>in Payload:[%+v], task msgId:[%+v] \n", string(task.Payload), task.ID)
				}
			},
		})
}

// create newAsynqClient 初使化任务 消费者
func newAsynqClient(c config.Config) *asynq.Client {
	return asynq.NewClient(
		asynq.RedisClientOpt{
			Addr:     c.Redis.Host,
			Password: c.Redis.Pass,
		})
}

//func newSchedulerbak(c config.Config) *asynq.AsynqScheduler {
//
//	location, _ := time.LoadLocation("Asia/Shanghai")
//	return asynq.NewScheduler(
//		asynq.RedisClientOpt{
//			Addr:     c.Redis.Host,
//			Password: c.Redis.Pass,
//		}, &asynq.SchedulerOpts{
//			Location: location,
//			EnqueueErrorHandler: func(task *asynq.Task, opts []asynq.Option, err error) {
//				fmt.Printf("AsynqScheduler EnqueueErrorHandler <<<<<<<===>>>>> err : %+v , task : %+v", err, task)
//			},
//		})
//}
