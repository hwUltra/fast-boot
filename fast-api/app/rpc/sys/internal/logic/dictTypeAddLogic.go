package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"github.com/jinzhu/copier"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictTypeAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictTypeAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictTypeAddLogic {
	return &DictTypeAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictTypeAddLogic) DictTypeAdd(in *sysPb.DictTypeForm) (*sysPb.SuccessResp, error) {
	item := model.SysDictTypeModel{}
	if err := copier.Copy(&item, in); err != nil {
		return nil, err
	}
	l.svcCtx.GormConn.Create(&item)
	return &sysPb.SuccessResp{}, nil
}
