package dept

import (
	"github.com/hwUltra/fb-tools/result"
	"net/http"

	"github.com/hwUltra/fb-tools/utils"
	"github.com/zeromicro/go-zero/rest/httpx"

	"fast-boot/app/api/admin/internal/logic/sys/dept"
	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"
)

// 修改部门
func UpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysDeptFormReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		if err := utils.ValidatorCheck(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := dept.NewUpdateLogic(r.Context(), svcCtx)
		resp, err := l.Update(&req)
		result.HttpResult(r, w, resp, err)
	}
}
