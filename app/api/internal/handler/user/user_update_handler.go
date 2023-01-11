package user

import (
	"net/http"

	"github.com/haimait/go-mindoc/app/api/internal/logic/user"
	"github.com/haimait/go-mindoc/app/api/internal/svc"
	"github.com/haimait/go-mindoc/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserUpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserUpdateLogic(r.Context(), svcCtx)
		resp, err := l.UserUpdate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
