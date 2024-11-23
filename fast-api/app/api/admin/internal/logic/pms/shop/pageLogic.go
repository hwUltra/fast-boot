package shop

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

// 店铺列表
func NewPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageLogic {
	return &PageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PageLogic) Page(req *types.ShopPageReq) (resp *types.ShopPageResp, err error) {
	res, err := l.svcCtx.PmsRpc.ShopPage(l.ctx, &pmsPb.PageReq{
		Keywords: req.Keywords,
		Status:   req.Status,
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
	})
	if err != nil {
		return nil, err
	}
	list := make([]*types.Shop, 0)
	_ = copier.Copy(&list, res.List)
	return &types.ShopPageResp{List: list, Total: res.Total}, nil
}
