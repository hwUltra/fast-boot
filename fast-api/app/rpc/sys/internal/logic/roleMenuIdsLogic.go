package logic

import (
	"context"
	"fast-boot/app/rpc/model"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleMenuIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleMenuIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleMenuIdsLogic {
	return &RoleMenuIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleMenuIdsLogic) RoleMenuIds(in *sysPb.IdReq) (*sysPb.RoleMenuIdsResp, error) {

	rmModel := model.SysRoleMenuModel{}
	menuIds := make([]int64, 0)
	l.svcCtx.GormClient.GormDb.Model(rmModel).
		Where("role_id = ?", in.Id).
		Pluck("menu_id", &menuIds)
	return &sysPb.RoleMenuIdsResp{Items: menuIds}, nil

}
