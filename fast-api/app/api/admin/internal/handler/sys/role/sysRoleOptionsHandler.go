package role

import (
	"fast-boot/common/result"
	"net/http"

	"fast-boot/app/api/admin/internal/logic/sys/role"
	"fast-boot/app/api/admin/internal/svc"
)

func SysRoleOptionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := role.NewSysRoleOptionsLogic(r.Context(), svcCtx)
		resp, err := l.SysRoleOptions()
		result.HttpResult(r, w, resp, err)
	}
}
