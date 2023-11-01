package other

import (
	"context"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

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

func (l *FileUploadLogic) FileUpload(req *types.FileUpload) (resp *types.FileUpload, err error) {
	return req, nil
}
