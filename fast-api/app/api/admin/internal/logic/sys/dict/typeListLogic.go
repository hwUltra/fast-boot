package dict

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TypeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TypeListLogic {
	return &TypeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TypeListLogic) TypeList(req *types.SysDictTypeListReq) (resp *types.SysDictTypeListResp, err error) {

	result, err := l.svcCtx.SysRpc.DictTypeList(l.ctx, &sysPb.DictTypeListReq{
		Keywords: req.Keywords,
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
	})
	if err != nil {
		return nil, err
	}

	list := make([]*types.SysDictTypeItem, 0)
	_ = copier.Copy(&list, result.List)

	return &types.SysDictTypeListResp{
		Total: result.Total,
		List:  list,
	}, nil
}
