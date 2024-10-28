package dict

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
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

func NewPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageLogic {
	return &PageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PageLogic) Page(req *types.SysDictPageReq) (resp *types.SysDictPageResp, err error) {
	result, err := l.svcCtx.SysRpc.DictPage(l.ctx, &sysPb.DictPageReq{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Keywords: req.Keywords})
	if err != nil {
		return nil, err
	}
	list := make([]*types.SysDictItem, 0)
	_ = copier.Copy(&list, result.List)

	return &types.SysDictPageResp{
		Total: result.Total,
		List:  list,
	}, nil
}
