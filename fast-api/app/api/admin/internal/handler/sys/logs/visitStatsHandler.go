package logs

import (
	"github.com/hwUltra/fb-tools/result"
	"net/http"

	"fast-boot/app/api/admin/internal/logic/sys/logs"
	"fast-boot/app/api/admin/internal/svc"
)

func VisitStatsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logs.NewVisitStatsLogic(r.Context(), svcCtx)
		resp, err := l.VisitStats()
		result.HttpResult(r, w, resp, err)
	}
}
