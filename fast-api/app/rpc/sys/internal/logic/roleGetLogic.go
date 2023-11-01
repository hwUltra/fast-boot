package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"github.com/jinzhu/copier"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleGetLogic {
	return &RoleGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RoleGet 获取单个
func (l *RoleGetLogic) RoleGet(in *sysPb.IdReq) (*sysPb.SysRole, error) {
	item := model.SysRoleModel{}
	l.svcCtx.GormConn.First(&item, in.Id)
	var sysRole sysPb.SysRole
	if err := copier.Copy(&sysRole, item); err != nil {
		return nil, err
	}
	return &sysRole, nil
}
