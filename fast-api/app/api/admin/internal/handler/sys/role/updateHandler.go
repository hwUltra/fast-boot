package role

import (
	"fast-boot/common/result"
	"net/http"

	"fast-boot/app/api/admin/internal/logic/sys/role"
	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysRoleFormReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := role.NewUpdateLogic(r.Context(), svcCtx)
		resp, err := l.Update(&req)
		result.HttpResult(r, w, resp, err)
	}
}
