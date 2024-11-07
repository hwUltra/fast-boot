package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeptOptionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptOptionsLogic {
	return &DeptOptionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptOptionsLogic) DeptOptions(in *sysPb.AnyReq) (*sysPb.DeptOptionResp, error) {
	deptModel := model.SysDeptModel{}
	var items []*model.SysDeptModel
	l.svcCtx.GormClient.GormDb.Model(deptModel).
		Where("status = 1").
		Order("sort asc, id asc ").
		Scan(&items)

	return &sysPb.DeptOptionResp{Items: GetTree(items, 0)}, nil
}

// GetTree 转化树结构
func GetTree(list []*model.SysDeptModel, parentID int64) []*sysPb.DeptOption {
	res := make([]*sysPb.DeptOption, 0)
	for _, item := range list {
		if item.ParentID == parentID {
			var dept sysPb.DeptOption
			dept.Label = item.Name
			dept.Value = item.Id
			dept.Children = GetTree(list, item.Id)
			res = append(res, &dept)
		}
	}
	return res
}
