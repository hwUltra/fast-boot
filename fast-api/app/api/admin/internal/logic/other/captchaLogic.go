package other

import (
	"context"
	captchaTool "github.com/hwUltra/fb-tools/captcha"

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
	var ct = captchaTool.NewCaptchaTool(captchaTool.CaptchaConf{
		Type:      captchaTool.MathType,
		Store:     captchaTool.RedisType,
		RedisConf: l.svcCtx.Config.Redis,
	})
	if id, b64s, _, err := ct.Make(); err != nil {
		return nil, err
	} else {
		//fmt.Printf(" id = %s \n b64s = %s \n answer = %s \n", id, b64s, answer)
		return &types.CaptchaResp{
			CaptchaBase64: b64s,
			CaptchaKey:    id,
		}, nil
	}

}
