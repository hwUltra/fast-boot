package dict

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysDictOptionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysDictOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDictOptionsLogic {
	return &SysDictOptionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysDictOptionsLogic) SysDictOptions() (resp []*types.SysDictOption, err error) {
	options, err := l.svcCtx.SysRpc.DictTypeOptions(l.ctx, &sysPb.AnyReq{})
	if err != nil {
		return nil, err
	}
	list := make([]*types.SysDictOption, 0)
	_ = copier.Copy(&list, options.Items)

	return list, nil
}
