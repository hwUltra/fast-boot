package logic

import (
	"context"
	"fast-boot/app/rpc/model"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleOptionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleOptionsLogic {
	return &RoleOptionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RoleOptions 下拉列表
func (l *RoleOptionsLogic) RoleOptions(in *sysPb.AnyReq) (*sysPb.RoleOptionsResp, error) {
	roleModel := model.SysRoleModel{}

	var items []*sysPb.RoleOption

	l.svcCtx.GormConn.Model(roleModel).
		Select("`id` as `value`", "`name` as `label` ").
		Where("`status` = 1").
		Order("`sort` asc, `id` asc ").
		Scan(&items)

	return &sysPb.RoleOptionsResp{Items: items}, nil
}
