package user

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.SysUserListReq) (resp *types.SysUserListResp, err error) {
	result, err := l.svcCtx.SysRpc.UserList(l.ctx, &sysPb.UserListReq{
		Keywords:  req.Keywords,
		DeptId:    req.DeptId,
		Status:    req.Status,
		PageSize:  req.PageSize,
		PageNum:   req.PageNum,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	})
	if err != nil {
		return nil, err
	}

	list := make([]*types.SysUserInfo, 0)
	_ = copier.Copy(&list, result.List)

	return &types.SysUserListResp{
		Total: result.Total,
		List:  list,
	}, nil
}
