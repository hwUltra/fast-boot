package goods

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 商品列表
func NewPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageLogic {
	return &PageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PageLogic) Page(req *types.PmsGoodsPageReq) (resp *types.PmsGoodsPageResp, err error) {
	res, err := l.svcCtx.PmsRpc.PmsGoodsPage(l.ctx, &pmsPb.PmsGoodsPageReq{
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

	list := make([]*types.PmsGoods, 0)
	_ = copier.Copy(&list, res.List)
	return &types.PmsGoodsPageResp{List: list, Total: res.Total}, nil
}
