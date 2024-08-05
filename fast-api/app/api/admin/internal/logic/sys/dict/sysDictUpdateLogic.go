package dict

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysDictUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysDictUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDictUpdateLogic {
	return &SysDictUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysDictUpdateLogic) SysDictUpdate(req *types.SysDictForm) error {
	dictForm := sysPb.DictForm{}
	_ = copier.Copy(&dictForm, req)

	if _, err := l.svcCtx.SysRpc.DictUpdate(l.ctx, &dictForm); err != nil {
		return nil
	}
	return nil
}
