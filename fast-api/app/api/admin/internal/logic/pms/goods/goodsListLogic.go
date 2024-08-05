package goods

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsListLogic {
	return &GoodsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsListLogic) GoodsList(req *types.PmsGoodsListReq) (resp *types.PmsGoodsListRsqp, err error) {
	res, err := l.svcCtx.PmsRpc.PmsGoodsList(l.ctx, &pmsPb.PmsGoodsListReq{
		ShopId:     req.ShopId,
		CategoryId: req.CategoryId,
		Keywords:   req.Keywords,
		Status:     req.Status,
		PageNum:    req.PageNum,
		PageSize:   req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	list := make([]types.PmsGoods, 0)
	_ = copier.Copy(&list, res.List)
	return &types.PmsGoodsListRsqp{List: list, Total: res.Total}, nil
}
