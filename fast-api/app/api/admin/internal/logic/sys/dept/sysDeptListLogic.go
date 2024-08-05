package dept

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysDeptListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysDeptListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDeptListLogic {
	return &SysDeptListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysDeptListLogic) SysDeptList(req *types.SysDeptListReq) (resp []types.SysDeptItem, err error) {
	items, err := l.svcCtx.SysRpc.DeptList(l.ctx, &sysPb.DeptListReq{Keywords: req.Keywords, Status: req.Status})
	if err != nil {
		return nil, err
	}
	list := make([]types.SysDeptItem, 0)
	_ = copier.Copy(&list, items.List)
	return list, nil
}
