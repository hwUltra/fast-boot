package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/app/rpc/pms/internal/svc"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/hwUltra/fb-tools/gormx"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShopDelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShopDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShopDelLogic {
	return &ShopDelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShopDelLogic) ShopDel(in *pmsPb.IdsReq) (*pmsPb.SuccessResp, error) {
	ct := (*model.PmsShopCache)(gormx.NewCacheTool(l.svcCtx.Config.Cache, l.svcCtx.GormClient.GormDb))
	ct.Del(in.Ids)
	return &pmsPb.SuccessResp{}, nil
}
