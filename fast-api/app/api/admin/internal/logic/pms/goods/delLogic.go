package goods

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除
func NewDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLogic {
	return &DelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelLogic) Del(req *types.PathIdsReq) (resp *types.NullResp, err error) {
	if _, err := l.svcCtx.PmsRpc.PmsGoodsDel(l.ctx, &pmsPb.IdsReq{Ids: req.Ids}); err != nil {
		return &types.NullResp{}, err
	}
	return &types.NullResp{}, nil
}
