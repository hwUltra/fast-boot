package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"strings"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuDelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuDelLogic {
	return &MenuDelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MenuDel 删除
func (l *MenuDelLogic) MenuDel(in *sysPb.IdsReq) (*sysPb.SuccessResp, error) {
	ids := strings.Split(in.Ids, ",")
	l.svcCtx.GormClient.GormDb.Delete(&model.SysMenuModel{}, ids)

	return &sysPb.SuccessResp{}, nil
}
