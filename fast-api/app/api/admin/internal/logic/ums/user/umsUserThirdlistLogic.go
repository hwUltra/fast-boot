package user

import (
	"context"
	"fast-boot/app/rpc/ums/umsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UmsUserThirdlistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUmsUserThirdlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UmsUserThirdlistLogic {
	return &UmsUserThirdlistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UmsUserThirdlistLogic) UmsUserThirdlist(req *types.UmsUserThirdlistReq) (resp *types.UserThirdListRsqp, err error) {
	result, err := l.svcCtx.UmsRpc.UserThirdList(l.ctx, &umsPb.ListReq{
		Keywords: req.Keywords,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	list := make([]*types.UserThird, 0)
	_ = copier.Copy(&list, result.List)

	return &types.UserThirdListRsqp{
		Total: result.Total,
		List:  list,
	}, nil
}
