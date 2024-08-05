package dept

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysDeptAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysDeptAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDeptAddLogic {
	return &SysDeptAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysDeptAddLogic) SysDeptAdd(req *types.SysDeptFormReq) (resp *types.SysDeptAddResp, err error) {
	deptForm := sysPb.DeptForm{}
	_ = copier.Copy(&deptForm, req)

	if _, err = l.svcCtx.SysRpc.DeptAdd(l.ctx, &deptForm); err != nil {
		return nil, err
	}

	return &types.SysDeptAddResp{
		Id: req.Id,
	}, nil
}
