package auth

import (
	"context"
	"fast-boot/app/api/wxapp/internal/svc"
	"fast-boot/app/api/wxapp/internal/types"
	"fast-boot/app/rpc/ums/umsPb"
	"github.com/zeromicro/go-zero/core/logx"
)

type BindLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBindLogic(ctx context.Context, svcCtx *svc.ServiceContext) BindLogic {
	return BindLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BindLogic) Bind(req types.BindReq) (resp *types.TokenResp, err error) {
	sessionKey := "SessionKey-" + req.OpenId
	sessionKeyVal, err := l.svcCtx.RedisClient.Get(sessionKey)
	if err != nil {
		return nil, err
	}

	plainData, err := l.svcCtx.WxMiniApp.Decrypt(sessionKeyVal, req.EncryptedData, req.Iv)
	phoneNumber := plainData.PhoneNumber

	res, err := l.svcCtx.UmsRpc.Bind(l.ctx, &umsPb.BindReq{
		OpenId: req.OpenId,
		Mobile: phoneNumber,
	})

	if err != nil {
		return nil, err
	}

	return &types.TokenResp{
		OpenID: req.OpenId,
		Token:  res.Token,
	}, nil
}
