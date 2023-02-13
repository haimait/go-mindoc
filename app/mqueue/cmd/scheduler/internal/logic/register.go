package logic

import (
	"context"
	"strconv"

	"go-mindoc/app/mqueue/cmd/scheduler/internal/svc"
)

type MqueueScheduler struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronScheduler(ctx context.Context, svcCtx *svc.ServiceContext) *MqueueScheduler {
	return &MqueueScheduler{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// RegisterCluster 注册product 即时发送消息，异步执行
func (l *MqueueScheduler) RegisterCluster() {
	// 发送十次
	for i := 0; i < 10; i++ {
		l.settleRecordCluster(strconv.Itoa(i))
	}
}

// RegisterScheduler 注册product 即时发送消息，异步执行
func (l *MqueueScheduler) RegisterScheduler() {
	l.settleRecordScheduler()
}
