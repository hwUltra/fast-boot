package dict

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TypeGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTypeGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TypeGetLogic {
	return &TypeGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TypeGetLogic) TypeGet(req *types.PathIdReq) (resp *types.SysDictTypeItem, err error) {
	res, err := l.svcCtx.SysRpc.DictTypeGet(l.ctx, &sysPb.IdReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var info types.SysDictTypeItem
	if err := copier.Copy(&info, res); err != nil {
		return nil, err
	}
	return &info, nil
}
