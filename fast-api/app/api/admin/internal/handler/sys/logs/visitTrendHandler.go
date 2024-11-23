package logs

import (
	"github.com/hwUltra/fb-tools/result"
	"net/http"

	"fast-boot/app/api/admin/internal/logic/sys/logs"
	"fast-boot/app/api/admin/internal/svc"
)

func VisitTrendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logs.NewVisitTrendLogic(r.Context(), svcCtx)
		resp, err := l.VisitTrend()
		result.HttpResult(r, w, resp, err)
	}
}
