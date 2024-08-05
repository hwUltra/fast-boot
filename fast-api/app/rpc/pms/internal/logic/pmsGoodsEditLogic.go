package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"gorm.io/gorm"
	"strconv"
	"strings"

	"fast-boot/app/rpc/pms/internal/svc"
	"fast-boot/app/rpc/pms/pmsPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PmsGoodsEditLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPmsGoodsEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PmsGoodsEditLogic {
	return &PmsGoodsEditLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PmsGoodsEditLogic) PmsGoodsEdit(in *pmsPb.PmsGoodsForm) (*pmsPb.SuccessIdResp, error) {

	//新产品
	m := model.PmsGoodsModel{}
	m.Id = in.Id
	m.Name = in.Name
	m.ShopId = in.ShopId
	m.CategoryId = in.CategoryId
	m.BrandId = in.BrandId
	m.OriginPrice = in.OriginPrice
	m.Price = in.Price
	m.Sales = 0
	m.PicUrl = in.PicUrl
	m.Album = strings.Join(in.SubPicUrls, ",")
	m.Unit = in.Unit
	m.Description = in.Description
	m.Detail = in.Detail

	if err := l.svcCtx.GormConn.Transaction(func(tx *gorm.DB) error {
		if in.Id == 0 {
			if err := tx.Create(&m).Error; err != nil {
				return err
			}
		} else {
			if err := tx.Updates(&m).Error; err != nil {
				return err
			}
		}
		if err := tx.Model(model.PmsGoodsAttributeModel{}).Where("goods_id = ?", m.Id).Update("status", 0).Error; err != nil {
			return err
		}
		//attrMap := make(map[string]int64)
		for _, attr := range in.AttrList {
			attrModel := model.PmsGoodsAttributeModel{}
			attrModel.GoodsId = m.Id
			attrModel.Name = attr.Name
			attrModel.Value = attr.Value
			attrModel.Type = 2
			attrModel.Id = attr.Id
			attrModel.Status = 1
			//if attr.Id == 0 {
			//	attrMap[attr.IdStr] = attr.Id
			//}
			if attr.Id == 0 {
				if err := tx.Create(&attrModel).Error; err != nil {
					return err
				}
			} else {
				if err := tx.Updates(&attrModel).Error; err != nil {
					return err
				}
			}
		}

		specMap := make(map[string]int64)
		specImgMap := make(map[int64]string)
		for _, spec := range in.SpecList {
			specModel := model.PmsGoodsAttributeModel{}
			specModel.GoodsId = m.Id
			specModel.Name = spec.Name
			specModel.PicUrl = spec.PicUrl
			specModel.Value = spec.Value
			specModel.Type = 1
			specModel.Id = spec.Id
			specModel.Status = 1
			if spec.Id == 0 {
				if err := tx.Create(&specModel).Error; err != nil {
					return err
				}
			} else {
				if err := tx.Updates(&specModel).Error; err != nil {
					return err
				}
			}
			if spec.Id == 0 {
				specMap[spec.IdStr] = specModel.Id
			}
			specImgMap[specModel.Id] = spec.PicUrl
		}

		if err := tx.Model(model.PmsSkuModel{}).Where("goods_id = ?", m.Id).Update("status", 0).Error; err != nil {
			return err
		}

		for _, sku := range in.SkuList {
			skuModel := model.PmsSkuModel{}
			skuModel.Id = sku.Id
			skuModel.GoodsId = m.Id
			skuModel.SkuSn = sku.SkuSn
			skuModel.Price = sku.Price
			skuModel.Stock = sku.Stock
			skuModel.LockedStock = sku.LockedStock
			skuModel.Name = sku.Name
			skuModel.Status = 1
			specIds := strings.Split(sku.SpecIds, "|")
			specIdArray := make([]string, 0)
			for _, specIdStr := range specIds {
				specId, err := strconv.Atoi(specIdStr)
				if err != nil {
					if val, ok := specMap[specIdStr]; ok {
						specId = int(val)
					}
				}
				if imgUrl, ok := specImgMap[int64(specId)]; ok {
					if len(imgUrl) > 0 {
						skuModel.PicURL = imgUrl
					}
				}
				specIdArray = append(specIdArray, strconv.Itoa(specId))
			}
			skuModel.SpecIds = strings.Join(specIdArray, "_")

			if sku.Id == 0 {
				if err := tx.Save(&skuModel).Error; err != nil {
					return err
				}
			} else {
				if err := tx.Updates(&skuModel).Error; err != nil {
					return err
				}
			}

		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &pmsPb.SuccessIdResp{Id: m.Id}, nil
}
