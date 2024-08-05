package menu

import (
	"fast-boot/common/result"
	"net/http"

	"fast-boot/app/api/admin/internal/logic/sys/menu"
	"fast-boot/app/api/admin/internal/svc"
)

func SysMenuOptionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := menu.NewSysMenuOptionsLogic(r.Context(), svcCtx)
		resp, err := l.SysMenuOptions()
		result.HttpResult(r, w, resp, err)
	}
}
