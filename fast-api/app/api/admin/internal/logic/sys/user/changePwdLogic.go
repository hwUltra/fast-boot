package user

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改用户密码
func NewChangePwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePwdLogic {
	return &ChangePwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePwdLogic) ChangePwd(req *types.SysUserChangePwdReq) (resp *types.NullResp, err error) {
	_, err = l.svcCtx.SysRpc.UserChangePwd(l.ctx, &sysPb.UserChangePwdReq{
		UserId:   req.UserId,
		Password: req.Password,
	})
	if err != nil {
		return &types.NullResp{}, err
	}

	return &types.NullResp{}, nil
}
