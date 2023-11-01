package auth

import (
	"net/http"

	"fast-boot/app/api/test/internal/logic/auth"
	"fast-boot/app/api/test/internal/svc"
	"fast-boot/app/api/test/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewTestLogic(r.Context(), svcCtx)
		resp, err := l.Test(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
