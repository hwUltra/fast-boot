package other

import (
	"github.com/hwUltra/fb-tools/result"
	"net/http"

	"fast-boot/app/api/admin/internal/logic/other"
	"fast-boot/app/api/admin/internal/svc"
)

// 验证码
func CaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := other.NewCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.Captcha()
		result.HttpResult(r, w, resp, err)
	}
}
