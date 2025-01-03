package dict

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DataGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDataGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DataGetLogic {
	return &DataGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DataGetLogic) DataGet(req *types.SysDictDataGetReq) (resp *types.SysDictDataItem, err error) {
	res, err := l.svcCtx.SysRpc.DictDataGet(l.ctx, &sysPb.IdReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var info types.SysDictDataItem
	if err := copier.Copy(&info, res); err != nil {
		return nil, err
	}
	return &info, nil
}
