package auth

import (
	"fast-boot/common/result"
	"net/http"

	"fast-boot/app/api/test/internal/logic/auth"
	"fast-boot/app/api/test/internal/svc"
)

func MeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := auth.NewMeLogic(r.Context(), svcCtx)
		resp, err := l.Me()
		result.HttpResult(r, w, resp, err)
	}
}
