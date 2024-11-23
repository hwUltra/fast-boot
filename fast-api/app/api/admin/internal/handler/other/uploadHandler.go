package other

import (
	"github.com/hwUltra/fb-tools/result"
	"net/http"

	"fast-boot/app/api/admin/internal/logic/other"
	"fast-boot/app/api/admin/internal/svc"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := other.NewUploadLogic(r.Context(), svcCtx)
		file, fileHeader, _ := r.FormFile("file")
		resp, err := l.Upload(file, fileHeader)
		result.HttpResult(r, w, resp, err)
	}
}
