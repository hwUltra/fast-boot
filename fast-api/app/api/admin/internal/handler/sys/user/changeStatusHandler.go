package user

import (
	"github.com/hwUltra/fb-tools/result"
	"net/http"

	"github.com/hwUltra/fb-tools/utils"
	"github.com/zeromicro/go-zero/rest/httpx"

	"fast-boot/app/api/admin/internal/logic/sys/user"
	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"
)

// 修改用户状态
func ChangeStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysUserChangeStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		if err := utils.ValidatorCheck(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewChangeStatusLogic(r.Context(), svcCtx)
		resp, err := l.ChangeStatus(&req)
		result.HttpResult(r, w, resp, err)
	}
}