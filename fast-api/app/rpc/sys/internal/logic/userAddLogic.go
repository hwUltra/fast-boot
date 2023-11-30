package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"
	"fast-boot/common/cryptx"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserAddLogic) UserAdd(in *sysPb.UserAddReq) (*sysPb.IdResp, error) {
	user := model.SysUserModel{}
	l.svcCtx.GormConn.Where("username = ?", in.Username).First(&user)
	if user.Id > 0 {
		return nil, status.Error(100, "该用户已存在")
	}
	if err := copier.Copy(&user, in); err != nil {
		return nil, err
	}
	if err := l.svcCtx.GormConn.Transaction(func(tx *gorm.DB) error {
		user.Password = cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
		tx.Create(&user)
		//tx.Delete(&model.SysUserRoleModel{}).Where("user_id = ?", user.Id)
		for _, roleId := range in.RoleIds {
			tx.FirstOrCreate(&model.SysUserRoleModel{}, model.SysUserRoleModel{
				RoleId: roleId,
				UserId: user.Id,
			})
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &sysPb.IdResp{}, nil
}
