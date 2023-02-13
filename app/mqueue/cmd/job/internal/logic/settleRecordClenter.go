package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"go-mindoc/app/mqueue/cmd/job/internal/svc"
	"go-mindoc/app/mqueue/cmd/job/jobtype"
)

// SettleRecordClusterHandler ...
type SettleRecordClusterHandler struct {
	svcCtx *svc.ServiceContext
}

func NewRecordClusterHandler(svcCtx *svc.ServiceContext) *SettleRecordClusterHandler {
	return &SettleRecordClusterHandler{
		svcCtx: svcCtx,
	}
}

// ProcessTask 消息者，具体处理逻辑 exec : if return err != nil , asynq will retry
func (l *SettleRecordClusterHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	//接收任务数据
	var in jobtype.DeliveryPayload
	if err := json.Unmarshal(t.Payload(), &in); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: SkipRetry:%v  payLoad:%+v \n", err, asynq.SkipRetry, t.Payload())
	}
	//逻辑处理start...
	fmt.Printf("consumer job demo SettleRecordClusterHandler -----> in:[%+v]  todo exec ....\n", in)

	return nil
}
