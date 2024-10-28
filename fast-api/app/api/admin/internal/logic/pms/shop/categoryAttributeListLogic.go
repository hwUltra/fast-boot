package shop

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryAttributeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 新增类型-属性
func NewCategoryAttributeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryAttributeListLogic {
	return &CategoryAttributeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryAttributeListLogic) CategoryAttributeList(req *types.PmsCategoryAttributeListReq) (resp []*types.PmsCategoryAttribute, err error) {
	res, err := l.svcCtx.PmsRpc.PmsCategoryAttributeList(l.ctx, &pmsPb.PmsCategoryAttributeListReq{
		CategoryId: req.CategoryId,
		Type:       req.Type,
	})
	if err != nil {
		return nil, err
	}
	list := make([]*types.PmsCategoryAttribute, 0)

	_ = copier.Copy(&list, res.List)
	return list, nil
}
