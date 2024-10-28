package shop

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 类型列表
func NewCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryListLogic {
	return &CategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryListLogic) CategoryList(req *types.PmsCategoryListReq) (resp []*types.PmsCategory, err error) {
	res, err := l.svcCtx.PmsRpc.PmsCategoryList(l.ctx, &pmsPb.PmsCategoryListReq{
		Keywords: req.Keywords,
		Status:   req.Status,
		ShopId:   req.ShopId,
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
	})
	if err != nil {
		return resp, err
	}
	list := make([]*types.PmsCategory, 0)
	_ = copier.Copy(&list, res.List)
	return list, nil
}
