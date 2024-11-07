package other

import (
	"context"
	"github.com/hwUltra/fb-tools/captchax"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 验证码
func NewCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CaptchaLogic {
	return &CaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CaptchaLogic) Captcha() (resp *types.CaptchaResp, err error) {
	var ct = captchax.NewCaptchaTool(captchax.CaptchaConf{
		Type:      captchax.MathType,
		Store:     captchax.RedisType,
		CacheConf: l.svcCtx.Config.CacheConf,
	})
	if id, b64s, _, err := ct.Make(); err != nil {
		return nil, err
	} else {
		return &types.CaptchaResp{
			CaptchaBase64: b64s,
			CaptchaKey:    id,
		}, nil
	}

}
