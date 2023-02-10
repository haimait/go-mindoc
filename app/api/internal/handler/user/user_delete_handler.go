package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-mindoc/app/api/internal/logic/user"
	"go-mindoc/app/api/internal/svc"
	"go-mindoc/app/api/internal/types"
)

func UserDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserDeleteRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserDeleteLogic(r.Context(), svcCtx)
		resp, err := l.UserDelete(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
