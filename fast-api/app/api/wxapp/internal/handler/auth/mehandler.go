package auth

import (
	"fast-boot/common/result"
	"net/http"

	"fast-boot/app/api/wxapp/internal/logic/auth"
	"fast-boot/app/api/wxapp/internal/svc"
)

func MeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := auth.NewMeLogic(r.Context(), svcCtx)
		resp, err := l.Me()
		result.HttpResult(r, w, resp, err)
	}
}
