package goods

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsGetLogic {
	return &GoodsGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsGetLogic) GoodsGet(req *types.GoodsGetReq) (resp *types.PmsGoods, err error) {
	res, err := l.svcCtx.PmsRpc.PmsGoodsGet(l.ctx, &pmsPb.IdReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var info types.PmsGoods
	if err := copier.Copy(&info, res); err != nil {
		return nil, err
	}

	return &info, nil
}
