package other

import (
	"github.com/hwUltra/fb-tools/result"
	"net/http"

	"fast-boot/app/api/admin/internal/logic/other"
	"fast-boot/app/api/admin/internal/svc"
)

// 文件上传
func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := other.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(r)
		result.HttpResult(r, w, resp, err)
	}
}
