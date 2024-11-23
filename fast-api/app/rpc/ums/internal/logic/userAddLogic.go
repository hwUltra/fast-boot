package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/app/rpc/ums/internal/svc"
	"fast-boot/app/rpc/ums/umsPb"
	"github.com/hwUltra/fb-tools/gormx"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserAddLogic) UserAdd(in *umsPb.UserForm) (*umsPb.SuccessIdResp, error) {
	item := model.UserModel{}
	if err := copier.Copy(&item, in); err != nil {
		return nil, err
	}
	ct := (*model.UserCache)(gormx.NewCacheTool(l.svcCtx.Config.Cache, l.svcCtx.GormClient.GormDb))
	if err := ct.Create(&item); err != nil {
		return nil, err
	}
	return &umsPb.SuccessIdResp{Id: item.Id}, nil
}
