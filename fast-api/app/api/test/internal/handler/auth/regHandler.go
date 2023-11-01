package auth

import (
	"fast-boot/common/result"
	"github.com/go-playground/validator/v10"
	"net/http"

	"fast-boot/app/api/test/internal/logic/auth"
	"fast-boot/app/api/test/internal/svc"
	"fast-boot/app/api/test/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		if err := validator.New().StructCtx(r.Context(), req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := auth.NewRegLogic(r.Context(), svcCtx)
		resp, err := l.Reg(&req)
		result.HttpResult(r, w, resp, err)
	}
}
