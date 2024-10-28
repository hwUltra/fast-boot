package dept

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OptionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 部门下拉列表
func NewOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OptionsLogic {
	return &OptionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OptionsLogic) Options() (resp []types.SysDeptOption, err error) {
	options, err := l.svcCtx.SysRpc.DeptOptions(l.ctx, &sysPb.AnyReq{})
	if err != nil {
		return nil, err
	}

	list := make([]types.SysDeptOption, 0)
	_ = copier.Copy(&list, options.Items)

	return list, nil
}