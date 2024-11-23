package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/app/rpc/ums/internal/svc"
	"fast-boot/app/rpc/ums/umsPb"
	"github.com/hwUltra/fb-tools/gormx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDelLogic {
	return &UserDelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserDelLogic) UserDel(in *umsPb.IdsReq) (*umsPb.SuccessResp, error) {

	ct := (*model.UserCache)(gormx.NewCacheTool(l.svcCtx.Config.Cache, l.svcCtx.GormClient.GormDb))
	ct.Del(in.Ids)
	return &umsPb.SuccessResp{}, nil

}
