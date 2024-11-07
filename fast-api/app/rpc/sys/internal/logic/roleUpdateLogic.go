package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"google.golang.org/grpc/status"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleUpdateLogic {
	return &RoleUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RoleUpdate 修改
func (l *RoleUpdateLogic) RoleUpdate(in *sysPb.RoleForm) (*sysPb.SuccessResp, error) {
	info := model.SysRoleModel{}
	l.svcCtx.GormClient.GormDb.First(&info, in.Id)

	if info.Id == 0 {
		logx.WithContext(l.ctx).Errorf("role不存在: %s", in.Id)
		return nil, status.Error(100, "role不存在")
	}
	if info.Name != in.Name {
		info.Name = in.Name
	}
	if info.Code != in.Code {
		info.Code = in.Code
	}
	if info.Sort != in.Sort {
		info.Sort = in.Sort
	}
	if info.Status != int8(in.Status) {
		info.Status = int8(in.Status)
	}
	if info.DataScope != int8(in.DataScope) {
		info.DataScope = int8(in.DataScope)
	}
	res := l.svcCtx.GormClient.GormDb.Save(&info)
	if res.Error != nil {
		return nil, res.Error
	}

	return &sysPb.SuccessResp{}, nil
}
