package other

import (
	"fast-boot/common/result"
	"net/http"

	"fast-boot/app/api/wxapp/internal/logic/other"
	"fast-boot/app/api/wxapp/internal/svc"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := other.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(r)
		result.HttpResult(r, w, resp, err)
	}
}
