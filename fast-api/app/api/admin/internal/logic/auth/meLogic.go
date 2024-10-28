package auth

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/hwUltra/fb-tools/jwtx"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取个人信息
func NewMeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MeLogic {
	return &MeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MeLogic) Me() (resp *types.MeResp, err error) {
	userId := jwtx.GetUid(l.ctx)
	userInfo, err := l.svcCtx.SysRpc.UserGet(l.ctx, &sysPb.IdReq{Id: userId})

	if err != nil {
		return nil, err
	}
	return &types.MeResp{
		UserId:   userId,
		Nickname: userInfo.Nickname,
		Username: userInfo.Username,
		Avatar:   userInfo.Avatar,
		Roles:    userInfo.Roles,
		Perms:    userInfo.Perms,
	}, nil
}
