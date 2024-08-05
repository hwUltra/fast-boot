package shop

import (
	"fast-boot/common/result"
	"net/http"

	"fast-boot/app/api/admin/internal/logic/pms/shop"
	"fast-boot/app/api/admin/internal/svc"
)

func ShopOptionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := shop.NewShopOptionsLogic(r.Context(), svcCtx)
		resp, err := l.ShopOptions()
		result.HttpResult(r, w, resp, err)
	}
}
