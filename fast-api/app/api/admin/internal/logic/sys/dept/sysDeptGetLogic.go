package dept

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysDeptGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysDeptGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDeptGetLogic {
	return &SysDeptGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysDeptGetLogic) SysDeptGet(req *types.SysDeptGetReq) (resp *types.SysDeptInfo, err error) {
	res, err := l.svcCtx.SysRpc.DeptGet(l.ctx, &sysPb.IdReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	var info types.SysDeptInfo
	if err := copier.Copy(&info, res); err != nil {
		return nil, err
	}

	return &info, nil
}
