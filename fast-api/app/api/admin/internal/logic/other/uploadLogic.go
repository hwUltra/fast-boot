package other

import (
	"context"
	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"
	"fast-boot/common/xerr"
	"fmt"
	"github.com/hwUltra/fb-tools/uploadx"
	"mime/multipart"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 文件上传
func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(file multipart.File, fileHeader *multipart.FileHeader) (resp *types.UploadInfo, err error) {
	oss := uploadx.NewOss(l.svcCtx.Config.OSSConf)
	res, err := oss.UploadFile(file, fileHeader)
	if err != nil {
		return nil, xerr.NewErrCodeMsg(10011, "请上传正确的文件")
	}
	return &types.UploadInfo{
		Name: res.Name,
		Url:  fmt.Sprintf("%s/%s", l.svcCtx.Config.OSSConf.MinIoConf.MinIOBasePath, res.Path),
	}, nil
}
