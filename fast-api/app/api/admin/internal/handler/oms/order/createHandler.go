package order

import (
	"github.com/hwUltra/fb-tools/result"
	"github.com/hwUltra/fb-tools/utils"
	"net/http"

	"fast-boot/app/api/admin/internal/logic/oms/order"
	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// CreateHandler 创建订单
func CreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderFormReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		if err := utils.ValidatorCheck(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := order.NewCreateLogic(r.Context(), svcCtx)
		resp, err := l.Create(&req)
		result.HttpResult(r, w, resp, err)
	}
}
