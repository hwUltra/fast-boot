package menu

import (
	"github.com/hwUltra/fb-tools/result"
	"net/http"

	"fast-boot/app/api/admin/internal/logic/sys/menu"
	"fast-boot/app/api/admin/internal/svc"
)

// 菜单下拉列表
func OptionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := menu.NewOptionsLogic(r.Context(), svcCtx)
		resp, err := l.Options()
		result.HttpResult(r, w, resp, err)
	}
}
