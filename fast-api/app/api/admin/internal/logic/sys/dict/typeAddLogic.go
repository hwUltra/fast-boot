package dict

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TypeAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTypeAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TypeAddLogic {
	return &TypeAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TypeAddLogic) TypeAdd(req *types.SysDictTypeForm) (resp *types.IdResp, err error) {
	dictTypeForm := sysPb.DictTypeForm{}
	_ = copier.Copy(&dictTypeForm, req)

	if _, err = l.svcCtx.SysRpc.DictTypeAdd(l.ctx, &dictTypeForm); err != nil {
		return nil, err
	}

	return &types.IdResp{
		Id: req.Id,
	}, nil
}
