package auth

import (
	"fast-boot/common/result"
	"net/http"

	"fast-boot/app/api/admin/internal/logic/auth"
	"fast-boot/app/api/admin/internal/svc"
)

func LogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := auth.NewLogoutLogic(r.Context(), svcCtx)
		err := l.Logout()
		result.HttpResult(r, w, nil, err)
	}
}
