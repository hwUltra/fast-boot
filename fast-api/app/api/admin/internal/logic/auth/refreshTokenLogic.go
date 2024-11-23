package auth

import (
	"context"
	"fast-boot/common/cachex"
	"fast-boot/common/globalkey"
	"fmt"
	"github.com/hwUltra/fb-tools/jwtx"
	"github.com/hwUltra/fb-tools/result"
	"time"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 刷新ToKen
func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken() (resp *types.TokenResp, err error) {
	userId := jwtx.GetUid(l.ctx)
	key := fmt.Sprintf(globalkey.CacheUserTokenKey, userId)
	cc := cachex.NewStore(l.svcCtx.Config.Cache)
	var res types.TokenResp
	if err := cc.Cache.Get(key, &res); err != nil {
		return nil, result.NewErrCodeMsg(10001, "token已过期")
	}
	now := time.Now().Unix()
	if now > res.RefreshAfter {
		accessExpire := l.svcCtx.Config.Auth.AccessExpire
		accessToken, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, userId)
		if err != nil {
			return nil, err
		}
		res = types.TokenResp{
			AccessToken:  accessToken,
			AccessExpire: now + accessExpire,
			RefreshAfter: now + accessExpire/2,
		}
		_ = cc.Cache.SetWithExpire(key, res, time.Duration(l.svcCtx.Config.Auth.AccessExpire)*time.Second)
	}
	return &res, nil

}
