package dept

import (
	"fast-boot/common/result"
	"net/http"

	"fast-boot/app/api/admin/internal/logic/sys/dept"
	"fast-boot/app/api/admin/internal/svc"
)

func SysDeptOptionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := dept.NewSysDeptOptionsLogic(r.Context(), svcCtx)
		resp, err := l.SysDeptOptions()
		result.HttpResult(r, w, resp, err)
	}
}
