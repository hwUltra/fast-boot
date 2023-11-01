package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"strings"

	"fast-boot/app/rpc/pms/internal/svc"
	"fast-boot/app/rpc/pms/pmsPb"

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
	ids := strings.Split(in.Ids, ",")
	l.svcCtx.GormConn.Delete(&model.PmsShopModel{}, ids)
	return &pmsPb.SuccessResp{}, nil
}
