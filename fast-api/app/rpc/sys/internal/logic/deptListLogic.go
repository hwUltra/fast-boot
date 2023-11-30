package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeptListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptListLogic {
	return &DeptListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptListLogic) DeptList(in *sysPb.DeptListReq) (*sysPb.DeptListResp, error) {
	deptModel := model.SysDeptModel{}
	var items []*model.SysDeptModel
	l.svcCtx.GormConn.Model(deptModel).
		Scopes(
			deptModel.WithStatus(in.Status),
			deptModel.WithKeywords(in.Keywords)).
		Order("sort asc, id asc ").
		Scan(&items)

	return &sysPb.DeptListResp{List: DeptListTree(items, 0)}, nil
}

// DeptListTree 转化树结构
func DeptListTree(list []*model.SysDeptModel, parentID int64) []*sysPb.SysDeptItem {
	res := make([]*sysPb.SysDeptItem, 0)
	for _, item := range list {
		if item.ParentID == parentID {
			var dept sysPb.SysDeptItem
			_ = copier.Copy(&dept, item)
			dept.Children = DeptListTree(list, item.Id)
			res = append(res, &dept)
		}
	}
	return res
}
