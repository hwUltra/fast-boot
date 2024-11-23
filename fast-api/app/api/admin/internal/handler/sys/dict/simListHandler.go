package dict

import (
	"github.com/hwUltra/fb-tools/result"
	"net/http"

	"fast-boot/app/api/admin/internal/logic/sys/dict"
	"fast-boot/app/api/admin/internal/svc"
)

func SimListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := dict.NewSimListLogic(r.Context(), svcCtx)
		resp, err := l.SimList()
		result.HttpResult(r, w, resp, err)
	}
}
