package auth

import (
	"context"
	"encoding/json"
	"fast-boot/app/rpc/sys/sysPb"
	"fast-boot/common/globalkey"
	"fast-boot/common/xerr"
	"fmt"
	captchaTool "github.com/hwUltra/fb-tools/captcha"
	"github.com/zeromicro/go-zero/core/stores/redis"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

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
	var ct = captchaTool.NewCaptchaTool(captchaTool.CaptchaConf{
		Type:      captchaTool.MathType,
		Store:     captchaTool.RedisType,
		RedisConf: l.svcCtx.Config.Redis,
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
	StoreUserInfo(res.UserId, resp, l.svcCtx.RedisClient, l.svcCtx.Config.Auth.AccessExpire)
	return resp, nil
}

func StoreUserInfo(userId int64, userInfo *types.TokenResp, redisClient *redis.Redis, expire int64) {
	bytes, _ := json.Marshal(userInfo)
	key := fmt.Sprintf(globalkey.CacheUserTokenKey, userId)
	_ = redisClient.Setex(key, string(bytes), int(expire))
}
