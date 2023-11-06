package other

import (
	"context"
	"fast-boot/common/miniox"
	"fast-boot/common/xerr"
	"fmt"
	"net/http"

	"fast-boot/app/api/wxapp/internal/svc"
	"fast-boot/app/api/wxapp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *http.Request) (resp *types.FileUploadResp, err error) {
	minioX := miniox.CreateMinioX(l.svcCtx.Config.MinioX)
	updateInfo, err := minioX.MinIOUpload(req)
	if err != nil {
		return nil, xerr.NewErrCodeMsg(10011, "请上传正确的文件")
	}
	return &types.FileUploadResp{
		Name: updateInfo.Name,
		Ext:  updateInfo.Ext,
		Size: updateInfo.Size,
		Hash: updateInfo.Hash,
		Path: fmt.Sprintf("%s/%s", l.svcCtx.Config.MinioX.MinIOBasePath, updateInfo.Path),
	}, nil

}
