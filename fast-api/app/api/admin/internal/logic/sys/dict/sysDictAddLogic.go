package dict

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysDictAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysDictAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDictAddLogic {
	return &SysDictAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysDictAddLogic) SysDictAdd(req *types.SysDictForm) (resp *types.SysDictAddResp, err error) {
	dictForm := sysPb.DictForm{}
	_ = copier.Copy(&dictForm, req)

	if _, err = l.svcCtx.SysRpc.DictAdd(l.ctx, &dictForm); err != nil {
		return nil, err
	}

	return &types.SysDictAddResp{
		Id: req.Id,
	}, nil
}
