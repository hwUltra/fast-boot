package user

import (
	"context"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeStatusLogic {
	return &ChangeStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeStatusLogic) ChangeStatus(req *types.SysUserChangeStatusReq) (err error) {
	// todo: add your logic here and delete this line

	return
}
