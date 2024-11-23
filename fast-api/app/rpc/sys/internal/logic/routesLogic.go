package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"github.com/hwUltra/fb-tools/gormx"
	"github.com/jinzhu/copier"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoutesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoutesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoutesLogic {
	return &RoutesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoutesLogic) Routes(in *sysPb.RoutesReq) (*sysPb.RoutesResp, error) {

	ct := (*model.SysMenuCache)(gormx.NewCacheTool(l.svcCtx.Config.Cache, l.svcCtx.GormClient.GormDb))
	items := ct.Routes(in.Uid)

	list := make([]*sysPb.SysMenu, 0)
	for _, item := range items {
		var menu sysPb.SysMenu
		_ = copier.Copy(&menu, item)

		roles := make([]string, 0)
		for _, v := range item.Roles {
			roles = append(roles, v.Code)
		}
		menu.Roles = roles

		list = append(list, &menu)
	}

	return &sysPb.RoutesResp{List: list}, nil
}
