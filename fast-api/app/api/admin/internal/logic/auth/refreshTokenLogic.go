package auth

import (
	"context"
	"encoding/json"
	"fast-boot/app/rpc/sys/sysPb"
	"fast-boot/common/globalkey"
	"fast-boot/common/xerr"
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
	exists, err := l.svcCtx.RedisClient.Exists(key)
	if exists {
		info, err := l.svcCtx.RedisClient.Get(key)
		if err == nil {
			var result *types.TokenResp
			if err = json.Unmarshal([]byte(info), &result); err != nil {
				return nil, xerr.NewErrMsg("token有误")
			}
			now := time.Now().Unix()
			if now > result.RefreshAfter {
				res, err := l.svcCtx.SysRpc.RefreshToken(l.ctx,
					&sysPb.RefreshTokenReq{UserId: userId})
				if err != nil {
					return nil, err
				}
				loginResp := &types.TokenResp{
					AccessToken:  res.AccessToken,
					AccessExpire: res.AccessExpire,
					RefreshAfter: res.RefreshAfter,
				}
				StoreUserInfo(userId, loginResp, l.svcCtx.RedisClient, l.svcCtx.Config.Auth.AccessExpire)
				return loginResp, nil
			} else {
				return result, nil
			}
		}

	}
	return nil, result.NewErrCodeMsg(10001, "token已过期")
}
