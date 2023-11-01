package goods

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditLogic {
	return &EditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditLogic) Edit(req *types.GoodsEditRep) (resp *types.IdResp, err error) {
	form := pmsPb.PmsGoodsForm{}
	_ = copier.Copy(&form, req)

	res, err := l.svcCtx.PmsRpc.PmsGoodsEdit(l.ctx, &form)
	if err != nil {
		return nil, err
	}
	return &types.IdResp{
		ID: res.Id,
	}, nil
}
