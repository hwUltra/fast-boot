package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"
	"fast-boot/common/cryptx"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserChangePwdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserChangePwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserChangePwdLogic {
	return &UserChangePwdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserChangePwdLogic) UserChangePwd(in *sysPb.UserChangePwdReq) (*sysPb.SuccessResp, error) {
	user := model.SysUserModel{}
	l.svcCtx.GormConn.First(&user, in.UserId)
	if user.Id == 0 {
		return nil, status.Error(100, "该用户不存在")
	}
	if len(in.Password) > 0 {
		user.Password = cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	}
	l.svcCtx.GormConn.Save(&user)
	return &sysPb.SuccessResp{}, nil
}
