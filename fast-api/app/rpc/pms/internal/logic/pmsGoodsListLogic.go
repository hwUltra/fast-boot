package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/common/xerr"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"fast-boot/app/rpc/pms/internal/svc"
	"fast-boot/app/rpc/pms/pmsPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PmsGoodsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPmsGoodsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PmsGoodsListLogic {
	return &PmsGoodsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PmsGoodsListLogic) PmsGoodsList(in *pmsPb.PmsGoodsListReq) (*pmsPb.PmsGoodsListResp, error) {
	m := model.PmsGoodsModel{}
	fmt.Println("PmsGoodsListReq", in)
	total := int64(0)
	if err := l.svcCtx.GormConn.Model(m).
		Scopes(
			m.WithShopId(in.ShopId),
			m.WithStatus(in.Status),
			m.WithKeywords(in.Keywords),
			m.WithCategoryId(in.CategoryId),
			m.WithCreatedAt(in.StartTime, in.EndTime)).
		Count(&total).Error; err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to get  err : %v , in :%+v", err, in)
	}

	list := make([]*pmsPb.PmsGoodsInfo, 0)

	if total > 0 {
		items := make([]*model.PmsGoodsModel, 0)
		l.svcCtx.GormConn.Model(m).
			Scopes(
				m.WithShopId(in.ShopId),
				m.WithStatus(in.Status),
				m.WithCategoryId(in.CategoryId),
				m.WithKeywords(in.Keywords),
				m.WithCreatedAt(in.StartTime, in.EndTime)).
			Preload("Category", "visible = 1").
			Preload("Brand").
			Preload("SkuList", "`status` = 1").
			Preload("SpecList", "`type` = 1 and `status` = 1").
			Preload("AttrList", "`type` = 2 and `status` = 1").
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

	return &pmsPb.PmsGoodsListResp{List: list, Total: total}, nil
}
