package dept

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysDeptDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysDeptDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDeptDelLogic {
	return &SysDeptDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysDeptDelLogic) SysDeptDel(req *types.SysDeptDelReq) error {
	if _, err := l.svcCtx.SysRpc.DeptDel(l.ctx, &sysPb.IdsReq{Ids: req.Ids}); err != nil {
		return err
	}
	return nil
}
