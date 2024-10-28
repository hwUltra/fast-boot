package dict

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Update
func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.SysDictForm) (resp *types.NullResp, err error) {
	form := sysPb.DictForm{}
	_ = copier.Copy(&form, req)

	if _, err := l.svcCtx.SysRpc.DictUpdate(l.ctx, &form); err != nil {
		return &types.NullResp{}, err
	}
	return &types.NullResp{}, nil
}
