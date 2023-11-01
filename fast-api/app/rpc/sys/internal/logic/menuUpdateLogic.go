package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"google.golang.org/grpc/status"
	"strconv"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuUpdateLogic {
	return &MenuUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MenuUpdate 修改
func (l *MenuUpdateLogic) MenuUpdate(in *sysPb.MenuForm) (*sysPb.SuccessResp, error) {
	info := model.SysMenuModel{}
	l.svcCtx.GormConn.First(&info, in.Id)

	if info.Id == 0 {
		logx.WithContext(l.ctx).Errorf("Menu不存在: %s", in.Id)
		return nil, status.Error(100, "Menu不存在")
	}

	info.ParentID = in.ParentId
	info.Name = in.Name
	info.Type = int8(in.Type)
	info.Path = in.Path
	info.Component = in.Component
	info.Perm = in.Perm
	info.Visible = int8(in.Visible)
	info.Sort = in.Sort
	info.Icon = in.Icon
	info.Redirect = in.Redirect

	treePath := "0"
	if in.ParentId > 0 {
		pTreePath := ""
		l.svcCtx.GormConn.Model(model.SysMenuModel{}).Where("id = ?", in.ParentId).Pluck("tree_path", &pTreePath)
		treePath = pTreePath + "," + strconv.Itoa(int(in.ParentId))
	}
	info.TreePath = treePath

	res := l.svcCtx.GormConn.Save(&info)
	if res.Error != nil {
		return nil, res.Error
	}

	return &sysPb.SuccessResp{}, nil
}
