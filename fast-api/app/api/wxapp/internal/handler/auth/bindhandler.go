package auth

import (
	"fast-boot/common/result"
	"net/http"

	"fast-boot/app/api/wxapp/internal/logic/auth"
	"fast-boot/app/api/wxapp/internal/svc"
	"fast-boot/app/api/wxapp/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func BindHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BindReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := auth.NewBindLogic(r.Context(), svcCtx)
		resp, err := l.Bind(req)
		result.HttpResult(r, w, resp, err)
	}
}
