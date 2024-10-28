package shop

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取店铺
func NewCategoryGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryGetLogic {
	return &CategoryGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryGetLogic) CategoryGet(req *types.PathIdReq) (resp *types.PmsCategory, err error) {
	res, err := l.svcCtx.PmsRpc.PmsCategoryGet(l.ctx, &pmsPb.IdReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var info types.PmsCategory
	if err := copier.Copy(&info, res); err != nil {
		return nil, err
	}

	return &info, nil
}
