package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuOptionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuOptionsLogic {
	return &MenuOptionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuOptionsLogic) MenuOptions(in *sysPb.AnyReq) (*sysPb.MenuOptionsResp, error) {
	menuModel := model.SysMenuModel{}
	var items []*model.SysMenuModel
	l.svcCtx.GormConn.Model(menuModel).
		Where("visible = true").
		Order("sort asc, id asc ").
		Scan(&items)

	return &sysPb.MenuOptionsResp{Items: GetTreeMenu(items, 0)}, nil
}

func GetTreeMenu(list []*model.SysMenuModel, parentID int64) []*sysPb.MenuOption {
	res := make([]*sysPb.MenuOption, 0)
	for _, item := range list {
		if item.ParentID == parentID {
			var dept sysPb.MenuOption
			dept.Label = item.Name
			dept.Value = item.Id
			dept.Children = GetTreeMenu(list, item.Id)
			res = append(res, &dept)
		}
	}
	return res
}
