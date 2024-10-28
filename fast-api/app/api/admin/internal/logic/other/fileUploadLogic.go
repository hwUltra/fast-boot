package other

import (
	"context"
	"fast-boot/common/xerr"
	"fmt"
	"github.com/hwUltra/fb-tools/miniox"
	"net/http"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 文件上传
func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(r *http.Request) (resp *types.FileUpload, err error) {
	minioX := miniox.CreateMinioX(l.svcCtx.Config.MinioX)
	res, err := minioX.MinIOUpload(r)
	if err != nil {
		return nil, xerr.NewErrCodeMsg(10011, "请上传正确的文件")
	}
	return &types.FileUpload{
		Name: res.Name,
		Ext:  res.Ext,
		Size: res.Size,
		Hash: res.Hash,
		Path: fmt.Sprintf("%s/%s", l.svcCtx.Config.MinioX.MinIOBasePath, res.Path),
	}, nil
}
