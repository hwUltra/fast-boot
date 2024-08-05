package dict

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysDictTypeAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysDictTypeAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDictTypeAddLogic {
	return &SysDictTypeAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysDictTypeAddLogic) SysDictTypeAdd(req *types.SysDictTypeForm) (resp *types.SysDictTypeAddResp, err error) {
	dictTypeForm := sysPb.DictTypeForm{}
	_ = copier.Copy(&dictTypeForm, req)

	if _, err = l.svcCtx.SysRpc.DictTypeAdd(l.ctx, &dictTypeForm); err != nil {
		return nil, err
	}

	return &types.SysDictTypeAddResp{
		Id: req.Id,
	}, nil
}
