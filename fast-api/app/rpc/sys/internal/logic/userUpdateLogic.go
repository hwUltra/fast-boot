package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"github.com/hwUltra/fb-tools/utils"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserUpdate 修改
func (l *UserUpdateLogic) UserUpdate(in *sysPb.UserUpdateReq) (*sysPb.SuccessResp, error) {
	user := model.SysUserModel{}
	l.svcCtx.GormClient.GormDb.First(&user, in.Id)
	if user.Id == 0 {
		return nil, status.Error(100, "该用户不存在")
	}
	//校验用户名
	if in.Username != "" && in.Username != user.Username {
		oldUser := model.SysUserModel{}
		l.svcCtx.GormClient.GormDb.Where("username = ?", in.Username).First(&oldUser)
		if oldUser.Id != 0 {
			return nil, status.Error(100, "用户名已存在")
		}
	}
	if len(in.Nickname) > 0 {
		user.Nickname = in.Nickname
	}
	if len(in.Username) > 0 {
		user.Username = in.Username
	}
	if len(in.Password) > 0 {
		user.Password = utils.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	}
	if in.DeptId > 0 {
		user.DeptID = in.DeptId
	}
	if len(in.Avatar) > 0 {
		user.Avatar = in.Avatar
	}
	if len(in.Mobile) > 0 {
		user.Mobile = in.Mobile
	}
	if len(in.Email) > 0 {
		user.Email = in.Email
	}
	if in.Gender != -1 {
		user.Gender = int8(in.Gender)
	}
	if in.Status != -1 {
		user.Status = int8(in.Status)
	}

	if err := l.svcCtx.GormClient.GormDb.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&user).Error; err != nil {
			return err
		}
		if err := tx.Where("user_id = ?", user.Id).Delete(&model.SysUserRoleModel{}).Error; err != nil {
			return err
		}
		for _, roleId := range in.RoleIds {
			if err := tx.FirstOrCreate(&model.SysUserRoleModel{}, model.SysUserRoleModel{
				RoleId: roleId,
				UserId: user.Id,
			}).Error; err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &sysPb.SuccessResp{}, nil
}
