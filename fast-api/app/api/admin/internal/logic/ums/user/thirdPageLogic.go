package user

import (
	"context"
	"fast-boot/app/rpc/ums/umsPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ThirdPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取第三方用户列表
func NewThirdPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThirdPageLogic {
	return &ThirdPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ThirdPageLogic) ThirdPage(req *types.UserPageReq) (resp *types.UserThirdPageRsqp, err error) {
	result, err := l.svcCtx.UmsRpc.UserThirdPage(l.ctx, &umsPb.PageReq{
		Keywords: req.Keywords,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	list := make([]*types.UserThird, 0)
	_ = copier.Copy(&list, result.List)

	return &types.UserThirdPageRsqp{
		Total: result.Total,
		List:  list,
	}, nil
}
