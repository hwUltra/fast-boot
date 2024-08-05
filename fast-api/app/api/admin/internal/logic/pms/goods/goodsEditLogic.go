package goods

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsEditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsEditLogic {
	return &GoodsEditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsEditLogic) GoodsEdit(req *types.GoodsEditRep) (resp *types.GoodsEditResp, err error) {
	form := pmsPb.PmsGoodsForm{}
	_ = copier.Copy(&form, req)

	res, err := l.svcCtx.PmsRpc.PmsGoodsEdit(l.ctx, &form)
	if err != nil {
		return nil, err
	}
	return &types.GoodsEditResp{
		Id: res.Id,
	}, nil
}
