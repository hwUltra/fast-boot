package user

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysUserDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysUserDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserDelLogic {
	return &SysUserDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysUserDelLogic) SysUserDel(req *types.SysUserDelReq) error {
	if _, err := l.svcCtx.SysRpc.UserDel(l.ctx, &sysPb.IdsReq{Ids: req.Ids}); err != nil {
		return err
	}
	return nil
}
