package other

import (
	"fast-boot/common/result"
	"net/http"

	"fast-boot/app/api/wxapp/internal/logic/other"
	"fast-boot/app/api/wxapp/internal/svc"
	"fast-boot/app/api/wxapp/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SendSmsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendSmsReq

		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := other.NewSendSmsLogic(r.Context(), svcCtx)
		resp, err := l.SendSms(req)
		result.HttpResult(r, w, resp, err)

		//if err := httpx.Parse(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}
		//
		//l := other.NewSendSmsLogic(r.Context(), svcCtx)
		//resp, err := l.SendSms(req)
		//if err != nil {
		//	httpx.Error(w, err)
		//} else {
		//	httpx.OkJson(w, resp)
		//}
	}
}
