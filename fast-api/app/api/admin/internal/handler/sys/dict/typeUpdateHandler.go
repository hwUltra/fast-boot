package dict

import (
	"fast-boot/common/result"
	"net/http"

	"fast-boot/app/api/admin/internal/logic/sys/dict"
	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TypeUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysDictTypeForm
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := dict.NewTypeUpdateLogic(r.Context(), svcCtx)
		err := l.TypeUpdate(&req)
		result.HttpResult(r, w, nil, err)
	}
}
