package menu

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysMenuGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysMenuGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysMenuGetLogic {
	return &SysMenuGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysMenuGetLogic) SysMenuGet(req *types.SysMenuGetReq) (resp *types.SysMenuInfo, err error) {
	res, err := l.svcCtx.SysRpc.MenuGet(l.ctx, &sysPb.IdReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var menuInfo types.SysMenuInfo
	if err := copier.Copy(&menuInfo, res); err != nil {
		return nil, err
	}

	return &menuInfo, nil
}
