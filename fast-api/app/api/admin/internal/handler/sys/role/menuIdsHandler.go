package role

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"github.com/hwUltra/fb-tools/result"
	"github.com/hwUltra/fb-tools/utils"

	"fast-boot/app/api/admin/internal/logic/sys/role"
	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"
)

// 获取角色的菜单ID集合
func MenuIdsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysMenuIdReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		if err := utils.ValidatorCheck(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		l := role.NewMenuIdsLogic(r.Context(), svcCtx)
		resp, err := l.MenuIds(&req)
		result.HttpResult(r, w, resp, err)
	}
}
