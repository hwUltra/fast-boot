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

type ShopGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShopGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShopGetLogic {
	return &ShopGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShopGetLogic) ShopGet(in *pmsPb.IdReq) (*pmsPb.Shop, error) {

	ct := (*model.PmsShopCache)(gormx.NewCacheTool(l.svcCtx.Config.Cache, l.svcCtx.GormClient.GormDb))
	item := ct.Get(in.Id)

	var info pmsPb.Shop
	if err := copier.Copy(&info, item); err != nil {
		return nil, err
	}

	return &info, nil
}
