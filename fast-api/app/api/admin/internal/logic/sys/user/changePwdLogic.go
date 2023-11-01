package user

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"fmt"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePwdLogic {
	return &ChangePwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePwdLogic) ChangePwd(req *types.SysUserChangePwdReq) (err error) {
	fmt.Println("UserChangePwdLogic", req)
	_, err = l.svcCtx.SysRpc.UserChangePwd(l.ctx, &sysPb.UserChangePwdReq{
		UserId:   req.UserId,
		Password: req.Password,
	})
	if err != nil {
		return err
	}

	return nil
}
