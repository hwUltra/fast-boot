package shop

import (
	"github.com/hwUltra/fb-tools/result"
	"net/http"

	"github.com/hwUltra/fb-tools/utils"
	"github.com/zeromicro/go-zero/rest/httpx"

	"fast-boot/app/api/admin/internal/logic/pms/shop"
	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"
)

// 删除
func DelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PathIdsReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		if err := utils.ValidatorCheck(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := shop.NewDelLogic(r.Context(), svcCtx)
		resp, err := l.Del(&req)
		result.HttpResult(r, w, resp, err)
	}
}