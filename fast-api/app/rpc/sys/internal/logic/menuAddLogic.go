package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"github.com/jinzhu/copier"
	"strconv"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuAddLogic {
	return &MenuAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MenuAdd 添加
func (l *MenuAddLogic) MenuAdd(in *sysPb.MenuForm) (*sysPb.SuccessResp, error) {
	item := &model.SysMenuModel{}
	if err := copier.Copy(&item, in); err != nil {
		return nil, err
	}
	//parentId - treePath
	treePath := "0"
	if in.ParentId > 0 {
		pTreePath := ""
		l.svcCtx.GormConn.Model(model.SysMenuModel{}).Where("id = ?", in.ParentId).Pluck("tree_path", &pTreePath)
		treePath = pTreePath + "," + strconv.Itoa(int(in.ParentId))
	}
	item.TreePath = treePath

	l.svcCtx.GormConn.Create(&item)
	return &sysPb.SuccessResp{}, nil
}
