package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/common/cryptx"
	"fast-boot/common/jwtx"
	"google.golang.org/grpc/status"
	"time"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *sysPb.LoginReq) (*sysPb.LoginResp, error) {
	res := &model.SysUserModel{}
	l.svcCtx.GormConn.Where("username = ?", in.Username).First(res)

	if res.Id == 0 {
		logx.WithContext(l.ctx).Errorf("用户不存在Username: %s", in.Username)
		return nil, status.Error(100, "用户不存在")
	}
	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if password != res.Password {
		return nil, status.Error(100, "密码错误")
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	accessToken, err := jwtx.GetToken(l.svcCtx.Config.JWT.AccessSecret, now, accessExpire, res.Id)
	if err != nil {
		return nil, err
	}
	return &sysPb.LoginResp{
		UserId:       res.Id,
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil

}
