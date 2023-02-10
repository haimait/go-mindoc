package ping

import (
	"go-mindoc/pkg/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-mindoc/app/api/internal/logic/ping"
	"go-mindoc/app/api/internal/svc"
	"go-mindoc/app/api/internal/types"
)

func PingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PingReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := ping.NewPingLogic(r.Context(), svcCtx)
		resp, err := l.Ping(&req)
		response.Response(w, resp, err) //②
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
	}
}
