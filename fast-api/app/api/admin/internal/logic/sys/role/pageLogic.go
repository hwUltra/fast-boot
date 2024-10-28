package role

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

// 角色分页列表
func NewPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PageLogic {
	return &PageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PageLogic) Page(req *types.SysRolePageReq) (resp *types.SysRolePageResp, err error) {
	result, err := l.svcCtx.SysRpc.RoleList(l.ctx, &sysPb.RoleListReq{
		Keywords: req.Keywords,
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
	})
	if err != nil {
		return nil, err
	}

	list := make([]*types.SysRoleInfo, 0)
	_ = copier.Copy(&list, result.List)

	return &types.SysRolePageResp{
		Total: result.Total,
		List:  list,
	}, nil
}
