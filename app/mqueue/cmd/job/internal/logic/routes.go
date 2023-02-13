package logic

import (
	"context"
	"github.com/hibiken/asynq"
	"go-mindoc/app/mqueue/cmd/job/internal/svc"
	"go-mindoc/app/mqueue/cmd/job/jobtype"
)

type CronJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronJob(ctx context.Context, svcCtx *svc.ServiceContext) *CronJob {
	return &CronJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Register  job 注册消费者Handle监听
func (l *CronJob) Register() *asynq.ServeMux {

	mux := asynq.NewServeMux()

	//ClusterSettleRecord 处理即时时任务Handler
	mux.Handle(jobtype.ClusterSettleRecord, NewRecordClusterHandler(l.svcCtx))

	//ScheduleSettleRecord scheduler job NewSettleRecordHandler 处理定时任务Handler
	mux.Handle(jobtype.ScheduleSettleRecord, NewSettleRecordShceduleHandler(l.svcCtx))

	//defer job
	//mux.Handle(jobtype.DeferCloseHomestayOrder, NewCloseHomestayOrderHandler(l.svcCtx))

	//queue job , asynq support queue job
	// wait you fill..

	return mux
}
