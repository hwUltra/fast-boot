package dict

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysDictListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysDictListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDictListLogic {
	return &SysDictListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysDictListLogic) SysDictList(req *types.SysDictListReq) (resp *types.SysDictListResp, err error) {
	result, err := l.svcCtx.SysRpc.DictList(l.ctx, &sysPb.DictListReq{
		TypeCode: req.TypeCode,
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
	})
	if err != nil {
		return nil, err
	}

	list := make([]*types.SysDictItem, 0)
	_ = copier.Copy(&list, result.List)

	return &types.SysDictListResp{
		Total: result.Total,
		List:  list,
	}, nil
}
