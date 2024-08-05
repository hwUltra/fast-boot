package dict

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysDictDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysDictDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDictDelLogic {
	return &SysDictDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysDictDelLogic) SysDictDel(req *types.SysDictDelReq) error {
	if _, err := l.svcCtx.SysRpc.DictDel(l.ctx, &sysPb.IdsReq{Ids: req.Ids}); err != nil {
		return err
	}
	return nil
}
