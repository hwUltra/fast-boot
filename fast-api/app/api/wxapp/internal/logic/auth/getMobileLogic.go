package auth

import (
	"context"
	"fast-boot/app/api/wxapp/internal/svc"
	"fast-boot/app/api/wxapp/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMobileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMobileLogic {
	return &GetMobileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMobileLogic) GetMobile(req *types.DecryptReq) (resp *types.MobileResp, err error) {
	sessionKey := "SessionKey-" + req.OpenId
	sessionKeyVal, err := l.svcCtx.RedisClient.Get(sessionKey)
	if err != nil {
		return nil, err
	}
	plainData, err := l.svcCtx.WxMiniApp.Decrypt(sessionKeyVal, req.EncryptedData, req.Iv)
	return &types.MobileResp{
		PhoneNumber:     plainData.PhoneNumber,
		PurePhoneNumber: plainData.PurePhoneNumber,
	}, nil
}
