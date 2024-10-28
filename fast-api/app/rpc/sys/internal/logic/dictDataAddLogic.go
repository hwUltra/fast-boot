package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"github.com/jinzhu/copier"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictDataAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictDataAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictDataAddLogic {
	return &DictDataAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictDataAddLogic) DictDataAdd(in *sysPb.DictDataForm) (*sysPb.SuccessResp, error) {
	item := model.SysDictDataModel{}
	if err := copier.Copy(&item, in); err != nil {
		return nil, err
	}
	l.svcCtx.GormConn.Create(&item)
	return &sysPb.SuccessResp{}, nil
}
