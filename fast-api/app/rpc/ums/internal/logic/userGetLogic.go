package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/app/rpc/ums/internal/svc"
	"fast-boot/app/rpc/ums/umsPb"
	"fast-boot/common/globalkey"
	"github.com/hwUltra/fb-tools/gormx"
	"github.com/jinzhu/copier"
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

	ct := (*model.UserCache)(gormx.NewCacheTool(l.svcCtx.Config.CacheConf, l.svcCtx.GormClient.GormDb))
	userInfo := ct.Get(in.Id)

	res := &umsPb.User{}
	if err := copier.CopyWithOption(&res, userInfo, globalkey.CopierProtoOptions); err != nil {
		return nil, err
	}

	return res, nil
}
