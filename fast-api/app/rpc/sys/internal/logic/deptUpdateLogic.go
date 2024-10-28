package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"google.golang.org/grpc/status"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptUpdateLogic {
	return &DeptUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptUpdateLogic) DeptUpdate(in *sysPb.DeptForm) (*sysPb.SuccessResp, error) {
	info := model.SysDeptModel{}
	l.svcCtx.GormConn.First(&info, in.Id)

	if info.Id == 0 {
		logx.WithContext(l.ctx).Errorf("不存在: %s", in.Id)
		return nil, status.Error(100, "不存在")
	}

	info.ParentID = in.ParentId
	info.TreePath = in.TreePath
	info.Name = in.Name
	info.Code = in.Code
	info.Sort = in.Sort
	info.Status = int8(in.Status)
	info.UpdateBy = in.UpdateBy

	res := l.svcCtx.GormConn.Save(&info)
	if res.Error != nil {
		return nil, res.Error
	}

	return &sysPb.SuccessResp{}, nil
}
