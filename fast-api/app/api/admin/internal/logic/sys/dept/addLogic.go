package dept

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

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.SysDeptFormReq) (resp *types.IdResp, err error) {
	deptForm := sysPb.DeptForm{}
	_ = copier.Copy(&deptForm, req)

	if _, err = l.svcCtx.SysRpc.DeptAdd(l.ctx, &deptForm); err != nil {
		return nil, err
	}

	return &types.IdResp{
		Id: req.Id,
	}, nil

}
