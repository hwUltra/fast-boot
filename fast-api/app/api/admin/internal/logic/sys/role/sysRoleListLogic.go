package role

import (
	"context"
	"fast-boot/app/rpc/sys/sysPb"
	"github.com/jinzhu/copier"

	"fast-boot/app/api/admin/internal/svc"
	"fast-boot/app/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysRoleListLogic {
	return &SysRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysRoleListLogic) SysRoleList(req *types.SysRoleListReq) (resp *types.SysRoleListResp, err error) {
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

	return &types.SysRoleListResp{
		Total: result.Total,
		List:  list,
	}, nil
}
