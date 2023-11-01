package dict

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TypeOptionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTypeOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TypeOptionsLogic {
	return &TypeOptionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TypeOptionsLogic) TypeOptions(req *types.SysTypeReq) (resp *types.SysDictOptionsResp, err error) {
	options, err := l.svcCtx.SysRpc.DictOptions(l.ctx, &sysPb.TypeReq{
		TypeCode: req.TypeCode,
	})
	if err != nil {
		return nil, err
	}
	list := make([]*types.SysDictOption, 0)
	_ = copier.Copy(&list, options.Items)

	return &types.SysDictOptionsResp{List: list}, nil
}
