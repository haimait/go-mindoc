package user

import (
	"net/http"

	"github.com/haimait/go-mindoc/pkg/response"
	"github.com/haimait/go-mindoc/user/api/internal/logic/user"
	"github.com/haimait/go-mindoc/user/api/internal/svc"
	"github.com/haimait/go-mindoc/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserLoginLogic(r.Context(), svcCtx)
		resp, err := l.UserLogin(&req)
		response.Response(w, resp, err) //②
		//if err != nil {
		//	//httpx.ErrorCtx(r.Context(), w, err)
		//	httpx.OkJsonCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
	}
}
