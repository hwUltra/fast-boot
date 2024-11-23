package dict

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLogic {
	return &DelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelLogic) Del(req *types.SysDictDelReq) (resp *types.SysDictDelResp, err error) {
	if _, err := l.svcCtx.SysRpc.DictDel(l.ctx, &sysPb.IdsReq{Ids: req.Ids}); err != nil {
		return nil, err
	}
	return &types.SysDictDelResp{}, nil
}
