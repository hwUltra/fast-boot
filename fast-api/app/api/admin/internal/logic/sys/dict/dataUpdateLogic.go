package dict

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DataUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDataUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DataUpdateLogic {
	return &DataUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DataUpdateLogic) DataUpdate(req *types.SysDictDataForm) (resp *types.NullResp, err error) {
	form := sysPb.DictDataForm{}
	_ = copier.Copy(&form, req)

	if _, err := l.svcCtx.SysRpc.DictDataUpdate(l.ctx, &form); err != nil {
		return &types.NullResp{}, err
	}
	return &types.NullResp{}, nil
}
