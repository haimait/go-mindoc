package logic

import (
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"go-mindoc/app/mqueue/cmd/job/jobtype"
	"log"
	"time"
)

// 即时发送 job 可以实时发送，调用一次，就发送一次------> go-zero-looklook/app/mqueue/cmd/job/internal/logic/settleRecord.go.
func (l *MqueueScheduler) settleRecordCluster(taskId string) {
	// step1 to mysql get data
	payload, err := json.Marshal(jobtype.DeliveryPayload{UserID: 1, TaskID: taskId})
	if err != nil {
		fmt.Printf("【settleRecordScheduler】 json.Marshal(DeliveryPayload failed : %v \n", taskId)
		return
	}
	// step2 init send data
	task := asynq.NewTask(jobtype.ClusterSettleRecord, payload)

	// step3 RegisterScheduler task
	// 延迟执行 exec
	//info, err := l.svcCtx.AsynqClient.Enqueue(task, asynq.ProcessIn(3*time.Second))
	// MaxRetry 重度次数 Timeout超时时间
	info, err := l.svcCtx.AsynqClient.Enqueue(task, asynq.MaxRetry(10), asynq.Timeout(3*time.Second))
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("【settleRecordCluster】 enqueued task: id=%s queue=%s", info.ID, info.Queue)
}

// 定时发送scheduler job 可定义循环执行 如每3s执行一次 ------> go-zero-looklook/app/mqueue/cmd/job/internal/logic/settleRecord.go.
func (l *MqueueScheduler) settleRecordScheduler() {
	// step1 to mysql get data
	taskID := "taskId:123"
	payload, err := json.Marshal(jobtype.DeliveryPayload{UserID: 1, TaskID: taskID})
	if err != nil {
		fmt.Printf("【settleRecordScheduler】 json.Marshal(DeliveryPayload failed : %v \n", taskID)
		return
	}
	// step2 init send data
	task := asynq.NewTask(jobtype.ScheduleSettleRecord, payload)

	// step3 RegisterScheduler task
	// exec
	//entryID, err := l.svcCtx.AsynqScheduler.RegisterScheduler("*/1 * * * *", task)
	entryID, err := l.svcCtx.AsynqScheduler.Register("@every 3s", task)
	//entryID, err := l.svcCtx.AsynqScheduler.RegisterScheduler("@every 3s", task)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("!!!MqueueSchedulerErr!!! ====> 【settleRecordScheduler】 registered  err:%+v , task:%+v \n", err, task)
	}
	fmt.Printf("【settleRecordScheduler】 registered an  entry: %q \n", entryID)
}
