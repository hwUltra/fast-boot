package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"github.com/jinzhu/copier"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictDataGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictDataGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictDataGetLogic {
	return &DictDataGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictDataGetLogic) DictDataGet(in *sysPb.IdReq) (*sysPb.SysDictData, error) {
	item := model.SysDictDataModel{}
	l.svcCtx.GormConn.First(&item, in.Id)
	var sysDictData sysPb.SysDictData
	if err := copier.Copy(&sysDictData, item); err != nil {
		return nil, err
	}
	return &sysDictData, nil
}
