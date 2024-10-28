package user

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户分页列表
func NewPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageLogic {
	return &PageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PageLogic) Page(req *types.SysUserPageReq) (resp *types.SysUserPageResp, err error) {
	result, err := l.svcCtx.SysRpc.UserPage(l.ctx, &sysPb.UserPageReq{
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

	return &types.SysUserPageResp{
		Total: result.Total,
		List:  list,
	}, nil
}
