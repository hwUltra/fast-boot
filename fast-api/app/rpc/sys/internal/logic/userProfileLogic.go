package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/common/globalkey"
	"google.golang.org/grpc/status"
	"strings"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserProfileLogic {
	return &UserProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserProfileLogic) UserProfile(in *sysPb.IdReq) (*sysPb.SysUserProfile, error) {
	gormDb := l.svcCtx.GormClient.GormDb
	userInfo := model.SysUserModel{}
	gormDb.Preload("Roles", "status = ?", 1).Preload("Dept", "status = 1").First(&userInfo, in.Id)

	if userInfo.Id == 0 {
		logx.WithContext(l.ctx).Errorf("用户不存在uid: %d", in.Id)
		return nil, status.Error(100, "用户不存在")
	}

	rolesName := make([]string, 0)
	for _, v := range userInfo.Roles {
		rolesName = append(rolesName, v.Name)
	}
	return &sysPb.SysUserProfile{
		Id:        userInfo.Id,
		Username:  userInfo.Username,
		Nickname:  userInfo.Nickname,
		Avatar:    userInfo.Avatar,
		Status:    int64(userInfo.Status),
		DeptId:    userInfo.DeptID,
		DeptName:  userInfo.Dept.Name,
		Email:     userInfo.Email,
		Gender:    int64(userInfo.Gender),
		Mobile:    userInfo.Mobile,
		CreatedAt: userInfo.CreatedAt.Format(globalkey.DateTimeFormatStandardTime),
		RoleNames: strings.Join(rolesName, ","),
	}, nil

}
