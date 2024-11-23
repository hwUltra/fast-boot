package auth

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"github.com/hwUltra/fb-tools/result"
	"github.com/hwUltra/fb-tools/utils"

	"fast-boot/app/api/admin/internal/logic/auth"
	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"
)

// 登录
func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		if err := utils.ValidatorCheck(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		l := auth.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		result.HttpResult(r, w, resp, err)
	}
}
