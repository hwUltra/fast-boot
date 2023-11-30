package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"github.com/jinzhu/copier"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleAddLogic {
	return &RoleAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RoleAdd 添加
func (l *RoleAddLogic) RoleAdd(in *sysPb.RoleForm) (*sysPb.IdResp, error) {
	item := model.SysRoleModel{}
	if err := copier.Copy(&item, in); err != nil {
		return nil, err
	}
	result := l.svcCtx.GormConn.Create(&item)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sysPb.IdResp{Id: item.Id}, nil
}
