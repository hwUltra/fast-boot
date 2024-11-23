package user

import (
	"context"
	"fast-boot/app/rpc/ums/umsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户列表
func NewPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageLogic {
	return &PageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PageLogic) Page(req *types.UserPageReq) (resp *types.UserPageRsqp, err error) {
	result, err := l.svcCtx.UmsRpc.UserPage(l.ctx, &umsPb.PageReq{
		Keywords: req.Keywords,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	list := make([]*types.UserInfo, 0)

	_ = copier.Copy(&list, result.List)

	return &types.UserPageRsqp{
		Total: result.Total,
		List:  list,
	}, nil
}
