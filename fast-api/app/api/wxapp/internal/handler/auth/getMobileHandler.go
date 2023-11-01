package auth

import (
	"fast-boot/common/result"
	"net/http"

	"fast-boot/app/api/wxapp/internal/logic/auth"
	"fast-boot/app/api/wxapp/internal/svc"
	"fast-boot/app/api/wxapp/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetMobileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DecryptReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := auth.NewGetMobileLogic(r.Context(), svcCtx)
		resp, err := l.GetMobile(&req)
		result.HttpResult(r, w, resp, err)
	}
}
