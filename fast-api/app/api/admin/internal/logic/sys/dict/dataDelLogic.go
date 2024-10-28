package dict

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DataDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDataDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DataDelLogic {
	return &DataDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DataDelLogic) DataDel(req *types.PathIdsReq) (resp *types.NullResp, err error) {
	if _, err := l.svcCtx.SysRpc.DictDataDel(l.ctx, &sysPb.IdsReq{Ids: req.Ids}); err != nil {
		return &types.NullResp{}, err
	}
	return &types.NullResp{}, nil
}
