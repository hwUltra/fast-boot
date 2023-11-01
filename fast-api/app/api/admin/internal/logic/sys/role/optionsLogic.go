package role

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

func NewOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OptionsLogic {
	return &OptionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OptionsLogic) Options() (resp *types.SysRoleOptionsResp, err error) {
	options, err := l.svcCtx.SysRpc.RoleOptions(l.ctx, &sysPb.AnyReq{})
	if err != nil {
		return nil, err
	}
	list := make([]*types.SysRoleOption, 0)
	_ = copier.Copy(&list, options.Items)

	return &types.SysRoleOptionsResp{List: list}, nil
}
