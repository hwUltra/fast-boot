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

type ShopAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShopAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShopAddLogic {
	return &ShopAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShopAddLogic) ShopAdd(in *pmsPb.ShopForm) (*pmsPb.SuccessIdResp, error) {
	item := model.PmsShopModel{}
	if err := copier.Copy(&item, in); err != nil {
		return nil, err
	}
	ct := (*model.PmsShopCache)(gormx.NewCacheTool(l.svcCtx.Config.CacheConf, l.svcCtx.GormClient.GormDb))
	if err := ct.Create(&item); err != nil {
		return nil, err
	}
	return &pmsPb.SuccessIdResp{Id: item.Id}, nil
}
