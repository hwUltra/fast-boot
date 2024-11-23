package auth

import (
	"context"
	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"
	"fast-boot/app/rpc/sys/sysPb"
	"fast-boot/common/cachex"
	"fast-boot/common/globalkey"
	"fast-boot/common/xerr"
	"github.com/hwUltra/fb-tools/captchax"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 登录
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.TokenResp, err error) {
	//checkCaptcha
	var ct = captchax.NewCaptchaTool(captchax.CaptchaConf{
		Type:      captchax.MathType,
		Store:     captchax.RedisType,
		CacheConf: l.svcCtx.Config.Cache,
	})
	if ok := ct.VerifyCaptcha(req.CaptchaKey, req.CaptchaCode, true); ok == false {
		return nil, xerr.NewErrMsg("验证码错误")
	}
	res, err := l.svcCtx.SysRpc.Login(l.ctx,
		&sysPb.LoginReq{Username: req.UserName, Password: req.Password})
	if err != nil {
		return nil, err
	}
	resp = &types.TokenResp{
		AccessToken:  res.AccessToken,
		RefreshAfter: res.RefreshAfter,
		AccessExpire: res.AccessExpire,
	}
	cc := cachex.NewStore(l.svcCtx.Config.Cache)
	_ = cc.Cache.SetWithExpire(cc.FormatPrimary(globalkey.CacheUserTokenKey, res.UserId), resp, time.Duration(l.svcCtx.Config.Auth.AccessExpire)*time.Second)
	return resp, nil
}
