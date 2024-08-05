package user

import (
	"fast-boot/common/result"
	"net/http"

	"fast-boot/app/api/admin/internal/logic/sys/user"
	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SysUserChangePwdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysUserChangePwdReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		l := user.NewSysUserChangePwdLogic(r.Context(), svcCtx)

		err := l.SysUserChangePwd(&req)
		result.HttpResult(r, w, nil, err)
	}
}
