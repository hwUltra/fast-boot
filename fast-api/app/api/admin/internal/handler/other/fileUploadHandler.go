package other

import (
	"fast-boot/common/miniox"
	"fast-boot/common/result"
	"fmt"
	"net/http"

	"fast-boot/app/api/admin/internal/logic/other"
	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUpload
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		minioX := miniox.CreateMinioX(svcCtx.Config.MinioX)
		updateInfo, err := minioX.MinIOUpload(r)
		if err != nil {
			result.HttpResult(r, w, nil, err)
			return
		}
		req.Name = updateInfo.Name
		req.Ext = updateInfo.Ext
		req.Size = updateInfo.Size
		req.Hash = updateInfo.Hash
		req.Path = fmt.Sprintf("%s/%s", svcCtx.Config.MinioX.MinIOBasePath, updateInfo.Path)

		l := other.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		result.HttpResult(r, w, resp, err)
	}
}
