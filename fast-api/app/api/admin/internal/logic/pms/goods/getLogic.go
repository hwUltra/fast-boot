package goods

import (
	"context"
	"fast-boot/app/rpc/pms/pmsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get(req *types.PathIdReq) (resp *types.PmsGoods, err error) {
	res, err := l.svcCtx.PmsRpc.PmsGoodsGet(l.ctx, &pmsPb.IdReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var info types.PmsGoods
	if err := copier.Copy(&info, res); err != nil {
		return nil, err
	}

	return &info, nil
}
