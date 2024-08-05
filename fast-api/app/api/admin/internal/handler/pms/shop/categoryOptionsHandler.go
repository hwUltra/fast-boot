package shop

import (
	"fast-boot/app/api/admin/internal/types"
	"fast-boot/common/result"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"fast-boot/app/api/admin/internal/logic/pms/shop"
	"fast-boot/app/api/admin/internal/svc"
)

func CategoryOptionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CategoryOptionsReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		l := shop.NewCategoryOptionsLogic(r.Context(), svcCtx)
		resp, err := l.CategoryOptions(&req)
		result.HttpResult(r, w, resp, err)
	}
}
