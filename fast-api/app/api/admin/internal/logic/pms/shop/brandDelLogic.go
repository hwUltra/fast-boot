package shop

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BrandDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除
func NewBrandDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandDelLogic {
	return &BrandDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BrandDelLogic) BrandDel(req *types.BrandDelReq) (resp *types.BrandDelResp, err error) {
	if _, err := l.svcCtx.PmsRpc.PmsBrandDel(l.ctx, &pmsPb.IdsReq{Ids: req.Ids}); err != nil {
		return nil, err
	}

	return &types.BrandDelResp{}, nil
}
