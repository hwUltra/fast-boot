package dict

import (
	"fast-boot/common/result"
	"net/http"

	"fast-boot/app/api/admin/internal/logic/sys/dict"
	"fast-boot/app/api/admin/internal/svc"
)

func SysDictOptionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := dict.NewSysDictOptionsLogic(r.Context(), svcCtx)
		resp, err := l.SysDictOptions()
		result.HttpResult(r, w, resp, err)
	}
}
