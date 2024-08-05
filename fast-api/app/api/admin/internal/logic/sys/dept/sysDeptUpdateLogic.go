package dept

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysDeptUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysDeptUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDeptUpdateLogic {
	return &SysDeptUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysDeptUpdateLogic) SysDeptUpdate(req *types.SysDeptFormReq) error {
	form := sysPb.DeptForm{}
	_ = copier.Copy(&form, req)

	if _, err := l.svcCtx.SysRpc.DeptUpdate(l.ctx, &form); err != nil {
		return err
	}

	return nil
}
