package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"go-mindoc/app/mqueue/cmd/job/internal/svc"
	"go-mindoc/app/mqueue/cmd/job/jobtype"
)

// SettleRecordShceduleHandler   shcedule billing to home business
type SettleRecordShceduleHandler struct {
	svcCtx *svc.ServiceContext
}

func NewSettleRecordShceduleHandler(svcCtx *svc.ServiceContext) *SettleRecordShceduleHandler {
	return &SettleRecordShceduleHandler{
		svcCtx: svcCtx,
	}
}

// ProcessTask 消息者，具体处理逻辑 exec : if return err != nil , asynq will retry
func (l *SettleRecordShceduleHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	//接收任务数据
	var in jobtype.DeliveryPayload
	if err := json.Unmarshal(t.Payload(), &in); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: SkipRetry:%v  payLoad:%+v \n", err, asynq.SkipRetry, t.Payload())
	}
	//逻辑处理start...
	fmt.Printf("shcedule consumer job demo NewSettleRecordHandler-----> in:[%+v]  todo exec ....\n", in)

	return nil
}
