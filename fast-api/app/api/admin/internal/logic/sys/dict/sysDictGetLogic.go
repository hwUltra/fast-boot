package dict

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysDictGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysDictGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDictGetLogic {
	return &SysDictGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysDictGetLogic) SysDictGet(req *types.SysDictGetReq) (resp *types.SysDictItem, err error) {
	res, err := l.svcCtx.SysRpc.DictGet(l.ctx, &sysPb.IdReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var info types.SysDictItem
	if err := copier.Copy(&info, res); err != nil {
		return nil, err
	}
	return &info, nil
}
