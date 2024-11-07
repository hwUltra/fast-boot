package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"google.golang.org/grpc/status"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserGetLogic {
	return &UserGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserGetLogic) UserGet(in *sysPb.IdReq) (*sysPb.UserGetResp, error) {

	gormDb := l.svcCtx.GormClient.GormDb
	userInfo := model.SysUserModel{}
	gormDb.Preload("Roles", "status = ?", 1).First(&userInfo, in.Id)

	if userInfo.Id == 0 {
		logx.WithContext(l.ctx).Errorf("用户不存在uid: %d", in.Id)
		return nil, status.Error(100, "用户不存在")
	}

	roleIds := make([]int64, 0)
	rolesCodes := make([]string, 0)
	for _, v := range userInfo.Roles {
		roleIds = append(roleIds, v.Id)
		rolesCodes = append(rolesCodes, v.Code)
	}
	//查询权限
	perms := make([]string, 0)
	gormDb.Model(&model.SysMenuModel{}).
		Joins("JOIN sys_role_menu on sys_menu.id = sys_role_menu.menu_id").
		Where("sys_role_menu.role_id IN ?", roleIds).
		Where("sys_menu.type = 4").
		Pluck("sys_menu.perm", &perms)

	return &sysPb.UserGetResp{
		Id:       userInfo.Id,
		Username: userInfo.Username,
		Nickname: userInfo.Nickname,
		Avatar:   userInfo.Avatar,
		Status:   int64(userInfo.Status),
		DeptId:   userInfo.DeptID,
		Email:    userInfo.Email,
		Gender:   int64(userInfo.Gender),
		Mobile:   userInfo.Mobile,
		RoleIds:  roleIds,
		Roles:    rolesCodes,
		Perms:    perms,
	}, nil

}
