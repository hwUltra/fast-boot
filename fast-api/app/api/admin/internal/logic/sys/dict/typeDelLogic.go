package dict

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TypeDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTypeDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TypeDelLogic {
	return &TypeDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TypeDelLogic) TypeDel(req *types.PathIdsReq) error {
	if _, err := l.svcCtx.SysRpc.DictTypeDel(l.ctx, &sysPb.IdsReq{Ids: req.Ids}); err != nil {
		return err
	}
	return nil
}
