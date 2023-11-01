package user

import (
	"context"
	"fast-boot/app/rpc/ums/umsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.ListReq) (resp *types.UserListRsqp, err error) {
	result, err := l.svcCtx.UmsRpc.UserList(l.ctx, &umsPb.ListReq{
		Keywords: req.Keywords,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	list := make([]*types.UserInfo, 0)
	_ = copier.Copy(&list, result.List)

	return &types.UserListRsqp{
		Total: result.Total,
		List:  list,
	}, nil
}
