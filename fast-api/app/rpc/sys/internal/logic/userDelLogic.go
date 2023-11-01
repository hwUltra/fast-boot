package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"strings"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

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

// UserDel 删除
func (l *UserDelLogic) UserDel(in *sysPb.IdsReq) (*sysPb.SuccessResp, error) {
	ids := strings.Split(in.Ids, ",")
	l.svcCtx.GormConn.Delete(&model.SysUserModel{}, ids)
	return &sysPb.SuccessResp{}, nil
}
