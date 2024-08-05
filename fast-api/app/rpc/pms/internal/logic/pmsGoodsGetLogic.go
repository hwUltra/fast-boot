package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/app/rpc/pms/internal/svc"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/jinzhu/copier"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type PmsGoodsGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPmsGoodsGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PmsGoodsGetLogic {
	return &PmsGoodsGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PmsGoodsGetLogic) PmsGoodsGet(in *pmsPb.IdReq) (*pmsPb.PmsGoodsInfo, error) {
	m := model.PmsGoodsModel{}
	l.svcCtx.GormConn.
		//Preload("Shop", "status = 1").
		Preload("Category", "visible = 1").
		Preload("Brand").
		Preload("SkuList", "status = 1").
		Preload("SpecList", "type = 1 and status = 1").
		Preload("AttrList", "type = 2 and status = 1").
		First(&m, in.Id)

	var info pmsPb.PmsGoodsInfo
	if err := copier.Copy(&info, m); err != nil {
		return nil, err
	}
	info.CategoryName = ""
	if m.Category.Id > 0 {
		info.CategoryName = m.Category.Name
	}
	info.BrandName = ""
	if m.Brand.Id > 0 {
		info.BrandName = m.Brand.Name
	}
	info.SubPicUrls = strings.Split(m.Album, ",")
	_ = copier.Copy(&info.SkuList, m.SkuList)
	_ = copier.Copy(&info.SpecList, m.SpecList)
	_ = copier.Copy(&info.AttrList, m.AttrList)

	return &info, nil
}
