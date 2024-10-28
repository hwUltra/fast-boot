package user

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"github.com/hwUltra/fb-tools/result"
	"github.com/hwUltra/fb-tools/utils"

	"fast-boot/app/api/admin/internal/logic/ums/user"
	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"
)

// 获取第三方用户列表
func ThirdPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		if err := utils.ValidatorCheck(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		l := user.NewThirdPageLogic(r.Context(), svcCtx)
		resp, err := l.ThirdPage(&req)
		result.HttpResult(r, w, resp, err)
	}
}