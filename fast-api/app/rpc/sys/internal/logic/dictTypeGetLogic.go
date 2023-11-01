package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"github.com/jinzhu/copier"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictTypeGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictTypeGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictTypeGetLogic {
	return &DictTypeGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictTypeGetLogic) DictTypeGet(in *sysPb.IdReq) (*sysPb.SysDictType, error) {
	item := model.SysDictTypeModel{}
	l.svcCtx.GormConn.First(&item, in.Id)
	var sysDeptType sysPb.SysDictType
	if err := copier.Copy(&sysDeptType, item); err != nil {
		return nil, err
	}
	return &sysDeptType, nil
}
