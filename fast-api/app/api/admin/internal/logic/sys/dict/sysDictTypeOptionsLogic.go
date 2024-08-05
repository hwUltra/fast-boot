package dict

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysDictTypeOptionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysDictTypeOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDictTypeOptionsLogic {
	return &SysDictTypeOptionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysDictTypeOptionsLogic) SysDictTypeOptions(req *types.SysTypeReq) ([]types.SysDictOption, error) {
	options, err := l.svcCtx.SysRpc.DictOptions(l.ctx, &sysPb.TypeReq{
		TypeCode: req.TypeCode,
	})
	if err != nil {
		return nil, err
	}
	list := make([]types.SysDictOption, 0)
	_ = copier.Copy(&list, options.Items)

	return list, nil
}
