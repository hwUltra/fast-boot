package user

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/hwUltra/fb-tools/jwtx"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProfileLogic {
	return &ProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Profile 获取用户
func (l *ProfileLogic) Profile() (resp *types.SysUserProfileResq, err error) {

	userId := jwtx.GetUid(l.ctx)
	userInfo, err := l.svcCtx.SysRpc.UserProfile(l.ctx, &sysPb.IdReq{Id: userId})
	if err != nil {
		return nil, err
	}
	return &types.SysUserProfileResq{
		Id:        userId,
		Mobile:    userInfo.Mobile,
		Nickname:  userInfo.Nickname,
		Username:  userInfo.Username,
		Avatar:    userInfo.Avatar,
		RoleNames: userInfo.RoleNames,
		DeptName:  userInfo.DeptName,
		CreatedAt: userInfo.CreatedAt,
		Email:     userInfo.Email,
		Gender:    userInfo.Gender,
	}, nil
}
