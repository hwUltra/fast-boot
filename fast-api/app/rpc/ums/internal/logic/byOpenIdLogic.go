package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/common/jwtx"
	"fmt"
	"time"

	"fast-boot/app/rpc/ums/internal/svc"
	"fast-boot/app/rpc/ums/umsPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ByOpenIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewByOpenIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ByOpenIdLogic {
	return &ByOpenIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ByOpenIdLogic) ByOpenId(in *umsPb.OpenIdReq) (*umsPb.UserInfoResp, error) {
	item := model.UserThirdModel{}
	l.svcCtx.GormConn.Where("`platform` = 'wxapp' and `openid` = ?", in.OpenId).
		Preload("User").First(&item)

	fmt.Println("item", item)

	token := ""
	if item.Id == 0 {
		//没有数据插入数据
		item.Openid = in.OpenId
		item.Platform = "wxapp"
		l.svcCtx.GormConn.Create(&item)
	} else {
		if item.Uid > 0 {
			now := time.Now().Unix()
			accessExpire := l.svcCtx.Config.JWT.AccessExpire
			token, _ = jwtx.GetToken(l.svcCtx.Config.JWT.AccessSecret, now, accessExpire, item.Uid)
			if item.User.Id > 0 {
				return &umsPb.UserInfoResp{
					Uid:      item.Uid,
					Nickname: item.User.Nickname,
					Avatar:   item.User.Avatar,
					Mobile:   item.User.Mobile,
					Token:    token,
				}, nil
			}

		}
	}
	return &umsPb.UserInfoResp{
		Uid: item.Uid,
	}, nil

}
