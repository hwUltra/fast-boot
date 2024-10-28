package logic

import (
	"context"
	"github.com/hwUltra/fb-tools/jwtx"
	"time"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RefreshToken 刷新token
func (l *RefreshTokenLogic) RefreshToken(in *sysPb.RefreshTokenReq) (*sysPb.LoginResp, error) {
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	accessToken, err := jwtx.GetToken(l.svcCtx.Config.JWT.AccessSecret, now, accessExpire, in.UserId)
	if err != nil {
		return nil, err
	}

	return &sysPb.LoginResp{
		UserId:       in.UserId,
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil

}
