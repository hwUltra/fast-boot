package dict

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SimListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// SimList
func NewSimListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SimListLogic {
	return &SimListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SimListLogic) SimList(req *types.NullReq) (resp []types.SysSimDict, err error) {
	res, err := l.svcCtx.SysRpc.DictSimList(l.ctx, &sysPb.AnyReq{})
	if err != nil {
		return nil, err
	}
	list := make([]types.SysSimDict, 0)
	_ = copier.Copy(&list, res.List)
	return list, nil
}
