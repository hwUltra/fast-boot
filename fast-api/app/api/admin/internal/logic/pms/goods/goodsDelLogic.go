package goods

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsDelLogic {
	return &GoodsDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsDelLogic) GoodsDel(req *types.GoodsDelReq) error {

	if _, err := l.svcCtx.PmsRpc.PmsGoodsDel(l.ctx, &pmsPb.IdsReq{Ids: req.Ids}); err != nil {
		return err
	}
	return nil

}
