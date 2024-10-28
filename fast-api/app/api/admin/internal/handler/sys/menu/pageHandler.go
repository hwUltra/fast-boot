package menu

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"github.com/hwUltra/fb-tools/result"
	"github.com/hwUltra/fb-tools/utils"

	"fast-boot/app/api/admin/internal/logic/sys/menu"
	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"
)

// 菜单列表
func PageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysMenuPageReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		if err := utils.ValidatorCheck(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		l := menu.NewPageLogic(r.Context(), svcCtx)
		resp, err := l.Page(&req)
		result.HttpResult(r, w, resp, err)
	}
}