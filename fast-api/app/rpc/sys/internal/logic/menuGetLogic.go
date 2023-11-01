package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"github.com/jinzhu/copier"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuGetLogic {
	return &MenuGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MenuGet 获取单个
func (l *MenuGetLogic) MenuGet(in *sysPb.IdReq) (*sysPb.SysMenu, error) {
	item := model.SysMenuModel{}
	l.svcCtx.GormConn.First(&item, in.Id)
	var sysRole sysPb.SysMenu
	if err := copier.Copy(&sysRole, item); err != nil {
		return nil, err
	}
	return &sysRole, nil
}
