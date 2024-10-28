package shop

import (
	"github.com/hwUltra/fb-tools/result"
	"net/http"

	"fast-boot/app/api/admin/internal/logic/pms/shop"
	"fast-boot/app/api/admin/internal/svc"
)

// 店铺下拉列表
func OptionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := shop.NewOptionsLogic(r.Context(), svcCtx)
		resp, err := l.Options()
		result.HttpResult(r, w, resp, err)
	}
}
