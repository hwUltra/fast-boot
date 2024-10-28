package shop

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryOptionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 类型下拉列表
func NewCategoryOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryOptionsLogic {
	return &CategoryOptionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryOptionsLogic) CategoryOptions(req *types.PathIdReq) (resp []*types.OptionItem, err error) {
	options, err := l.svcCtx.PmsRpc.PmsCategoryOptions(l.ctx, &pmsPb.IdReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	list := make([]*types.OptionItem, 0)
	_ = copier.Copy(&list, options.Items)

	return list, nil
}
