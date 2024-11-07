package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"github.com/jinzhu/copier"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptGetLogic {
	return &DeptGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptGetLogic) DeptGet(in *sysPb.IdReq) (*sysPb.SysDept, error) {
	item := model.SysDeptModel{}
	l.svcCtx.GormClient.GormDb.First(&item, in.Id)
	var sysDept sysPb.SysDept
	if err := copier.Copy(&sysDept, item); err != nil {
		return nil, err
	}
	return &sysDept, nil
}
