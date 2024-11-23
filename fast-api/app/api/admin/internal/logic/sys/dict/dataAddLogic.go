package dict

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DataAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDataAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DataAddLogic {
	return &DataAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DataAddLogic) DataAdd(req *types.SysDictDataForm) (resp *types.SysDictDataFormResp, err error) {
	form := sysPb.DictDataForm{}
	_ = copier.Copy(&form, req)

	if _, err = l.svcCtx.SysRpc.DictDataAdd(l.ctx, &form); err != nil {
		return nil, err
	}

	return &types.SysDictDataFormResp{
		Id: req.Id,
	}, nil
}
