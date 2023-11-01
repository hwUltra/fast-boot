package logic

import (
	"context"
	"fast-boot/app/rpc/model"

	"fast-boot/app/rpc/pms/internal/svc"
	"fast-boot/app/rpc/pms/pmsPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PmsBrandOptionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPmsBrandOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PmsBrandOptionsLogic {
	return &PmsBrandOptionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PmsBrandOptionsLogic) PmsBrandOptions(in *pmsPb.IdReq) (*pmsPb.OptionsResp, error) {
	m := model.PmsBrandModel{}

	var items []*pmsPb.OptionItem

	l.svcCtx.GormConn.Model(m).
		Scopes(
			m.WithShopId(in.Id)).
		Select("`id` as `value`", "`name` as `label` ").
		Order("`sort` asc ,`id` asc ").
		Scan(&items)

	return &pmsPb.OptionsResp{Items: items}, nil
}
