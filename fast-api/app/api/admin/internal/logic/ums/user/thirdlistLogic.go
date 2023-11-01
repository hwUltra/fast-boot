package user

import (
	"context"
	"fast-boot/app/rpc/ums/umsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ThirdlistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewThirdlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThirdlistLogic {
	return &ThirdlistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ThirdlistLogic) Thirdlist(req *types.ListReq) (resp *types.UserThirdListRsqp, err error) {
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
