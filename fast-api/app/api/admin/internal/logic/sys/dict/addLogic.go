package dict

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewAddLogic Add
func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.SysDictForm) (resp *types.SysDictFormResp, err error) {
	dictForm := sysPb.DictForm{}
	_ = copier.Copy(&dictForm, req)

	if _, err = l.svcCtx.SysRpc.DictAdd(l.ctx, &dictForm); err != nil {
		return nil, err
	}

	return &types.SysDictFormResp{
		Id: req.Id,
	}, nil
}
