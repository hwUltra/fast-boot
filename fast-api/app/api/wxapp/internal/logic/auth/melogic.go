package auth

import (
	"context"
	"fast-boot/app/api/wxapp/internal/svc"
	"fast-boot/app/api/wxapp/internal/types"
	"fast-boot/app/rpc/ums/umsPb"
	"github.com/hwUltra/fb-tools/jwtx"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type MeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMeLogic(ctx context.Context, svcCtx *svc.ServiceContext) MeLogic {
	return MeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MeLogic) Me() (resp *types.UserInfo, err error) {
	uid := jwtx.GetUid(l.ctx)
	res, err := l.svcCtx.UmsRpc.UserGet(l.ctx, &umsPb.IdReq{Id: uid})
	if err != nil {
		return nil, err
	}
	userInfo := types.UserInfo{}
	if err := copier.Copy(&userInfo, res); err != nil {
		return nil, err
	}
	return &userInfo, err
}
