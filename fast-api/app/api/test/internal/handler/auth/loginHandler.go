package auth

import (
	"fast-boot/common/result"
	"net/http"

	"fast-boot/app/api/test/internal/logic/auth"
	"fast-boot/app/api/test/internal/svc"
	"fast-boot/app/api/test/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := auth.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		result.HttpResult(r, w, resp, err)
	}
}
