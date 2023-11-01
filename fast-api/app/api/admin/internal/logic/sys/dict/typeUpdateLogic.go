package dict

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TypeUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTypeUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TypeUpdateLogic {
	return &TypeUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TypeUpdateLogic) TypeUpdate(req *types.SysDictTypeForm) error {
	dictTypeForm := sysPb.DictTypeForm{}
	_ = copier.Copy(&dictTypeForm, req)

	if _, err := l.svcCtx.SysRpc.DictTypeUpdate(l.ctx, &dictTypeForm); err != nil {
		return nil
	}
	return nil
}
