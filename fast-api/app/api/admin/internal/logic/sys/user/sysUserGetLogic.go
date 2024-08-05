package user

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysUserGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysUserGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserGetLogic {
	return &SysUserGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysUserGetLogic) SysUserGet(req *types.SysUserGetReq) (resp *types.SysUserGet, err error) {
	res, err := l.svcCtx.SysRpc.UserGet(l.ctx, &sysPb.IdReq{Id: req.Id})

	if err != nil {
		return nil, err
	}

	userInfo := types.SysUserGet{}
	_ = copier.Copy(&userInfo, res)

	return &userInfo, nil
}
