package shop

import (
	"fast-boot/common/result"
	"net/http"

	"fast-boot/app/api/admin/internal/logic/pms/shop"
	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CategoryAttributeAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PmsCategoryAttributeForm
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := shop.NewCategoryAttributeAddLogic(r.Context(), svcCtx)
		err := l.CategoryAttributeAdd(&req)
		result.HttpResult(r, w, nil, err)
	}
}
