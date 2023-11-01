package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"github.com/jinzhu/copier"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictGetLogic {
	return &DictGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictGetLogic) DictGet(in *sysPb.IdReq) (*sysPb.SysDict, error) {
	item := model.SysDictModel{}
	l.svcCtx.GormConn.First(&item, in.Id)
	var sysDept sysPb.SysDict
	if err := copier.Copy(&sysDept, item); err != nil {
		return nil, err
	}
	return &sysDept, nil
}
