package shop

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"github.com/hwUltra/fb-tools/result"
	"github.com/hwUltra/fb-tools/utils"

	"fast-boot/app/api/admin/internal/logic/pms/shop"
	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"
)

// 类型下拉列表
func CategoryOptionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PmsCategoryOptionsReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		if err := utils.ValidatorCheck(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		l := shop.NewCategoryOptionsLogic(r.Context(), svcCtx)
		resp, err := l.CategoryOptions(&req)
		result.HttpResult(r, w, resp, err)
	}
}
