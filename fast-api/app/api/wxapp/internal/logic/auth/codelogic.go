package auth

import (
	"context"
	"fast-boot/app/api/wxapp/internal/svc"
	"fast-boot/app/api/wxapp/internal/types"
	"fast-boot/app/rpc/ums/umsPb"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type CodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) CodeLogic {
	return CodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CodeLogic) Code(req types.AuthCodeReq) (resp *types.TokenResp, err error) {

	code2Session, err := l.svcCtx.WxMiniApp.Code2Session(req.Code)
	if err != nil {
		return nil, err
	}
	sessionKey := "SessionKey-" + code2Session.OpenID
	err = l.svcCtx.RedisClient.Setex(sessionKey, code2Session.SessionKey, 48*60*60)
	if err != nil {
		return nil, err
	}
	fmt.Println("code2Session", code2Session)
	res, err := l.svcCtx.UmsRpc.ByOpenId(l.ctx, &umsPb.OpenIdReq{OpenId: code2Session.OpenID})
	if err != nil {
		return nil, err
	}

	return &types.TokenResp{
		OpenID: code2Session.OpenID,
		Token:  res.Token,
	}, nil

}
