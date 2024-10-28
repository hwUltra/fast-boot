package dict

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DataPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDataPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DataPageLogic {
	return &DataPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DataPageLogic) DataPage(req *types.SysDictDataPageReq) (resp *types.SysDictDataPageResp, err error) {
	result, err := l.svcCtx.SysRpc.DictDataPage(l.ctx, &sysPb.DictDataPageReq{
		DictId:   req.DictId,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Keywords: req.Keywords})
	if err != nil {
		return nil, err
	}
	list := make([]*types.SysDictDataItem, 0)
	_ = copier.Copy(&list, result.List)

	return &types.SysDictDataPageResp{
		Total: result.Total,
		List:  list,
	}, nil
}
