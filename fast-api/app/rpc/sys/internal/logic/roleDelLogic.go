package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"strings"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleDelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDelLogic {
	return &RoleDelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RoleDel 删除
func (l *RoleDelLogic) RoleDel(in *sysPb.IdsReq) (*sysPb.SuccessResp, error) {

	ids := strings.Split(in.Ids, ",")
	l.svcCtx.GormClient.GormDb.Delete(&model.SysRoleModel{}, ids)

	return &sysPb.SuccessResp{}, nil
}
