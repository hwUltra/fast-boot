package callBack

import (
	"fast-boot/app/api/wxapp/internal/logic/callBack"
	"github.com/hwUltra/fb-tools/result"
	"net/http"

	"fast-boot/app/api/wxapp/internal/svc"
)

func WxPayHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := callBack.NewWxPayLogic(r.Context(), svcCtx)
		resp, err := l.WxPay(r)
		result.HttpResult(r, w, resp, err)
		return
	}
}
