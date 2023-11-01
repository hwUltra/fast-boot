package dept

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.SysDeptListReq) (resp *types.SysDeptListResp, err error) {
	items, err := l.svcCtx.SysRpc.DeptList(l.ctx, &sysPb.DeptListReq{Keywords: req.Keywords, Status: req.Status})
	if err != nil {
		return nil, err
	}
	list := make([]*types.SysDeptItem, 0)
	_ = copier.Copy(&list, items.List)
	return &types.SysDeptListResp{List: list}, nil

}
