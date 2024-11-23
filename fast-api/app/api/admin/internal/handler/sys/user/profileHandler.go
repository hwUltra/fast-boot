package user

import (
	"github.com/hwUltra/fb-tools/result"
	"net/http"

	"fast-boot/app/api/admin/internal/logic/sys/user"
	"fast-boot/app/api/admin/internal/svc"
)

// Profile 获取用户
func ProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewProfileLogic(r.Context(), svcCtx)
		resp, err := l.Profile()
		result.HttpResult(r, w, resp, err)
	}
}
