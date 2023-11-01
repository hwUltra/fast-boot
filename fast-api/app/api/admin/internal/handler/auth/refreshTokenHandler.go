package auth

import (
	"fast-boot/common/result"
	"net/http"

	"fast-boot/app/api/admin/internal/logic/auth"
	"fast-boot/app/api/admin/internal/svc"
)

func RefreshTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := auth.NewRefreshTokenLogic(r.Context(), svcCtx)
		resp, err := l.RefreshToken()
		result.HttpResult(r, w, resp, err)
	}
}
