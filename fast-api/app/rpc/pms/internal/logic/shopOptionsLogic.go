package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/app/rpc/pms/internal/svc"
	"fast-boot/app/rpc/pms/pmsPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShopOptionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShopOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShopOptionsLogic {
	return &ShopOptionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShopOptionsLogic) ShopOptions(in *pmsPb.AnyReq) (*pmsPb.OptionsResp, error) {
	shopModel := model.PmsShopModel{}

	var items []*pmsPb.OptionItem

	l.svcCtx.GormConn.Model(shopModel).
		Select("id as value", "name as label ").
		Where("status = 1").
		Order("id asc ").
		Scan(&items)

	return &pmsPb.OptionsResp{Items: items}, nil
}
