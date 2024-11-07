package other

import (
	"fast-boot/app/api/admin/internal/logic/other"
	"fast-boot/app/api/admin/internal/svc"
	"github.com/hwUltra/fb-tools/result"
	"net/http"
)

// 文件上传
func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := other.NewUploadLogic(r.Context(), svcCtx)
		file, fileHeader, _ := r.FormFile("file")
		resp, err := l.Upload(file, fileHeader)
		result.HttpResult(r, w, resp, err)
	}
}
