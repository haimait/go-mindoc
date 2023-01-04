package user

import (
	"net/http"

	"github.com/haimait/go-mindoc/user/api/internal/logic/user"
	"github.com/haimait/go-mindoc/user/api/internal/svc"
	"github.com/haimait/go-mindoc/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpDateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserUpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUpDateUserLogic(r.Context(), svcCtx)
		resp, err := l.UpDateUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
