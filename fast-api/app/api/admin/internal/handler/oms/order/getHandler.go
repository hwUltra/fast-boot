package order

import (
	"github.com/hwUltra/fb-tools/result"
	"net/http"

	"github.com/hwUltra/fb-tools/utils"
	"github.com/zeromicro/go-zero/rest/httpx"

	"fast-boot/app/api/admin/internal/logic/oms/order"
	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"
)

// 获取订单
func GetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PathIdReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		if err := utils.ValidatorCheck(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := order.NewGetLogic(r.Context(), svcCtx)
		resp, err := l.Get(&req)
		result.HttpResult(r, w, resp, err)
	}
}
