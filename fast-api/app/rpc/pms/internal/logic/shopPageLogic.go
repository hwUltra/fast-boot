package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"github.com/hwUltra/fb-tools/gormx"
	"github.com/jinzhu/copier"

	"fast-boot/app/rpc/pms/internal/svc"
	"fast-boot/app/rpc/pms/pmsPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShopPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShopPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShopPageLogic {
	return &ShopPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShopPageLogic) ShopPage(in *pmsPb.PageReq) (*pmsPb.ShopPageResp, error) {
	shopModel := model.PmsShopModel{}

	total := int64(0)
	if err := l.svcCtx.GormClient.GormDb.Model(shopModel).
		Scopes(
			shopModel.WithStatus(in.Status),
			shopModel.WithKeywords(in.Keywords)).
		Count(&total).Error; err != nil {
		return nil, err
	}

	list := make([]*pmsPb.Shop, 0)

	if total > 0 {
		items := make([]*model.PmsShopModel, 0)
		l.svcCtx.GormClient.GormDb.Model(shopModel).
			Scopes(
				gormx.Paginate(int(in.PageNum), int(in.PageSize)),
				shopModel.WithStatus(in.Status),
				shopModel.WithKeywords(in.Keywords)).
			Order("id asc").Find(&items)

		for _, item := range items {
			var it pmsPb.Shop
			_ = copier.Copy(&it, item)

			list = append(list, &it)
		}
	}

	return &pmsPb.ShopPageResp{List: list, Total: total}, nil
}
