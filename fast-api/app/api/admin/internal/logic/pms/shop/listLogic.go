package shop

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"
	"fmt"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.ListReq) (resp *types.ShopListResp, err error) {
	res, err := l.svcCtx.PmsRpc.ShopList(l.ctx, &pmsPb.ListReq{
		Keywords: req.Keywords,
		Status:   req.Status,
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
	})
	fmt.Println("ListLogic", res)
	if err != nil {
		return nil, err
	}
	list := make([]*types.Shop, 0)
	_ = copier.Copy(&list, res.List)
	return &types.ShopListResp{List: list, Total: res.Total}, nil
}
