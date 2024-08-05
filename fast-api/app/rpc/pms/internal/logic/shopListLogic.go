package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/common/gormV2"
	"fast-boot/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"fast-boot/app/rpc/pms/internal/svc"
	"fast-boot/app/rpc/pms/pmsPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShopListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShopListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShopListLogic {
	return &ShopListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShopListLogic) ShopList(in *pmsPb.ListReq) (*pmsPb.ShopListResp, error) {
	shopModel := model.PmsShopModel{}

	total := int64(0)
	if err := l.svcCtx.GormConn.Model(shopModel).
		Scopes(
			shopModel.WithStatus(in.Status),
			shopModel.WithKeywords(in.Keywords)).
		Count(&total).Error; err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to get  err : %v , in :%+v", err, in)
	}

	list := make([]*pmsPb.Shop, 0)

	if total > 0 {
		items := make([]*model.PmsShopModel, 0)
		l.svcCtx.GormConn.Model(shopModel).
			Scopes(
				gormV2.Paginate(int(in.PageNum), int(in.PageSize)),
				shopModel.WithStatus(in.Status),
				shopModel.WithKeywords(in.Keywords)).
			Order("id asc").Find(&items)

		for _, item := range items {
			var it pmsPb.Shop
			_ = copier.Copy(&it, item)

			list = append(list, &it)
		}
	}

	return &pmsPb.ShopListResp{List: list, Total: total}, nil
}
