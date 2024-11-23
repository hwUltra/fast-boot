package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/app/rpc/pms/internal/svc"
	"fast-boot/app/rpc/pms/pmsPb"
	"fmt"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type PmsGoodsPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPmsGoodsPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PmsGoodsPageLogic {
	return &PmsGoodsPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PmsGoodsPageLogic) PmsGoodsPage(in *pmsPb.PmsGoodsPageReq) (*pmsPb.PmsGoodsPageResp, error) {
	m := model.PmsGoodsModel{}
	fmt.Println("PmsGoodsListReq", in)
	total := int64(0)
	if err := l.svcCtx.GormClient.GormDb.Model(m).
		Scopes(
			m.WithShopId(in.ShopId),
			m.WithStatus(in.Status),
			m.WithKeywords(in.Keywords),
			m.WithCategoryId(in.CategoryId),
			m.WithCreatedAt(in.StartTime, in.EndTime)).
		Count(&total).Error; err != nil {
		return nil, err
	}

	list := make([]*pmsPb.PmsGoodsInfo, 0)

	if total > 0 {
		items := make([]*model.PmsGoodsModel, 0)
		l.svcCtx.GormClient.GormDb.Model(m).
			Scopes(
				m.WithShopId(in.ShopId),
				m.WithStatus(in.Status),
				m.WithCategoryId(in.CategoryId),
				m.WithKeywords(in.Keywords),
				m.WithCreatedAt(in.StartTime, in.EndTime)).
			Preload("Category", "visible = true").
			Preload("Brand").
			Preload("SkuList", "status = 1").
			Preload("SpecList", "type = 1 and status = 1").
			Preload("AttrList", "type = 2 and status = 1").
			Order("id asc").Find(&items)

		for _, item := range items {
			var info pmsPb.PmsGoodsInfo
			if err := copier.Copy(&info, item); err != nil {
				return nil, err
			}
			info.CategoryName = ""
			if item.Category.Id > 0 {
				info.CategoryName = item.Category.Name
			}
			info.BrandName = ""
			if item.Brand.Id > 0 {
				info.BrandName = item.Brand.Name
			}
			_ = copier.Copy(&info.SkuList, item.SkuList)
			_ = copier.Copy(&info.SpecList, item.SpecList)
			_ = copier.Copy(&info.AttrList, item.AttrList)
			list = append(list, &info)
		}
	}

	return &pmsPb.PmsGoodsPageResp{List: list, Total: total}, nil
}
