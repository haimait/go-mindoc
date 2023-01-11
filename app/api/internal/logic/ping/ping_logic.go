package ping

import (
	"context"
	"github.com/haimait/go-mindoc/app/api/internal/svc"
	"github.com/haimait/go-mindoc/app/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//func (l *PingLogic) Ping(req *types.PingReq) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var resp *types.PingResp
//		resp.Msg = "pong"
//		response.Response(w, resp, nil) //②
//
//	}
//}

func (l *PingLogic) Ping(req *types.PingReq) (resp *types.PingResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.PingResp)
	resp.Msg = "pong"
	return resp, nil
}
