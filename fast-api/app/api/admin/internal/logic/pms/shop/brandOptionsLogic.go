package shop

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BrandOptionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBrandOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandOptionsLogic {
	return &BrandOptionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BrandOptionsLogic) BrandOptions(req *types.PathIdReq) (resp *types.OptionsResp, err error) {
	options, err := l.svcCtx.PmsRpc.PmsBrandOptions(l.ctx, &pmsPb.IdReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	list := make([]*types.OptionItem, 0)
	_ = copier.Copy(&list, options.Items)

	return &types.OptionsResp{List: list}, nil
}
