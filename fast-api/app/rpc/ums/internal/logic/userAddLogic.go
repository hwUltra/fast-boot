package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"fast-boot/app/rpc/ums/internal/svc"
	"fast-boot/app/rpc/ums/umsPb"

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
		return nil, errors.Wrapf(xerr.NewErrMsg("数据转化有误"), "UserAdd err:%v", err)
	}

	l.svcCtx.GormConn.Create(&item)

	return &umsPb.SuccessIdResp{}, nil
}
