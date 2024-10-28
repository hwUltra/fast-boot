package shop

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 新增类型
func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.ShopForm) (resp *types.IdResp, err error) {
	form := pmsPb.ShopForm{}
	_ = copier.Copy(&form, req)
	if _, err = l.svcCtx.PmsRpc.ShopAdd(l.ctx, &form); err != nil {
		return nil, err
	}

	return &types.IdResp{
		Id: req.Id,
	}, nil
}
