package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/app/rpc/ums/internal/svc"
	"fast-boot/app/rpc/ums/umsPb"
	"fast-boot/common/globalkey"
	"fast-boot/common/xerr"
	"github.com/pkg/errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserUpdate  修改用户
func (l *UserUpdateLogic) UserUpdate(in *umsPb.UserForm) (*umsPb.SuccessResp, error) {
	info := model.UserModel{}
	l.svcCtx.GormConn.Where("id = ?", in.Id).First(&info)
	if info.Id == 0 {
		return nil, errors.Wrapf(xerr.NewErrMsg("用户不存在"), "用户不存在：%d ", in.Id)
	}

	if len(in.Username) > 0 && (in.Username != info.Username) {
		info.Username = in.Username
	}
	if len(in.Nickname) > 0 && (in.Nickname != info.Nickname) {
		info.Nickname = in.Nickname
	}
	if len(in.Mobile) > 0 && (in.Mobile != info.Mobile) {
		info.Mobile = in.Mobile
	}
	if len(in.Avatar) > 0 && (in.Avatar != info.Avatar) {
		info.Avatar = in.Avatar
	}
	if len(in.IdCard) > 0 && (in.IdCard != info.IdCard) {
		info.IdCard = in.IdCard
	}
	if len(in.Signature) > 0 && (in.Signature != info.Signature) {
		info.Signature = in.Signature
	}
	if len(in.Birthday) > 0 {
		//info.Birthday = in.Birthday
		t, _ := time.ParseInLocation(globalkey.DateTimeFormatYMDTime, in.Birthday, time.Local) //这里按照当前时区转
		info.Birthday = t
	}
	if len(in.Tags) > 0 && (in.Tags != info.Tags) {
		info.Tags = in.Tags
	}
	if in.Gender >= 0 {
		info.Gender = int8(in.Gender)
	}
	if len(in.Source) > 0 && (in.Source != info.Source) {
		info.Source = in.Source
	}
	if in.SourceUid > 0 {
		info.SourceUid = in.SourceUid
	}

	res := l.svcCtx.GormConn.Save(&info)
	if res.Error != nil {
		return nil, res.Error
	}
	return &umsPb.SuccessResp{}, nil

}
