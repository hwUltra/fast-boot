package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"github.com/jinzhu/copier"

	"fast-boot/app/rpc/ums/internal/svc"
	"fast-boot/app/rpc/ums/umsPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserGetLogic {
	return &UserGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserGetLogic) UserGet(in *umsPb.IdReq) (*umsPb.User, error) {
	userInfo := model.UserModel{}
	l.svcCtx.GormConn.Where("`id` = ?", in.Id).First(&userInfo)

	res := &umsPb.User{}
	if err := copier.Copy(&res, userInfo); err != nil {
		return nil, err
	}

	//特别处理
	res.Birthday = userInfo.Birthday.Time.Format("2006-01-02")
	return res, nil
}
