package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/app/rpc/ums/internal/svc"
	"fast-boot/app/rpc/ums/umsPb"
	"fast-boot/common/jwtx"
	"gorm.io/gorm"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindLogic {
	return &BindLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BindLogic) Bind(in *umsPb.BindReq) (*umsPb.UserInfoResp, error) {
	userInfo := model.UserModel{}
	userThirdInfo := model.UserThirdModel{}
	l.svcCtx.GormConn.Where("`mobile` = ?", in.Mobile).First(&userInfo)
	l.svcCtx.GormConn.Where("`platform` = 'wxapp' and `openid` = ?", in.OpenId).
		First(&userThirdInfo)

	if userInfo.Id == 0 {
		//注册用户
		userInfo.Mobile = in.Mobile
		if err := l.svcCtx.GormConn.Transaction(func(tx *gorm.DB) error {
			if err := tx.Create(&userInfo).Error; err != nil {
				return err
			}
			if userThirdInfo.Uid == 0 {
				userThirdInfo.Uid = userInfo.Id
				if err := tx.Save(&userThirdInfo).Error; err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			return nil, err
		}
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	token, _ := jwtx.GetToken(l.svcCtx.Config.JWT.AccessSecret, now, accessExpire, userInfo.Id)

	return &umsPb.UserInfoResp{
		Uid:      userInfo.Id,
		Nickname: userInfo.Nickname,
		Avatar:   userInfo.Avatar,
		Mobile:   userInfo.Mobile,
		Token:    token,
	}, nil
}
