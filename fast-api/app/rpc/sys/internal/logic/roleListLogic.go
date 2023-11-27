package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/common/gormV2"
	"fast-boot/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RoleList  角色
func (l *RoleListLogic) RoleList(in *sysPb.RoleListReq) (*sysPb.RoleListResp, error) {
	roleModel := model.SysRoleModel{}

	total := int64(0)
	if err := l.svcCtx.GormConn.Model(roleModel).
		Scopes(roleModel.WithKeywords(in.Keywords)).
		Count(&total).Error; err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to get  err : %v , in :%+v", err, in)
	}
	list := make([]*sysPb.SysRole, 0)
	if total > 0 {
		items := make([]*model.SysRoleModel, 0)
		l.svcCtx.GormConn.Model(roleModel).
			Scopes(
				gormV2.Paginate(int(in.PageNum), int(in.PageSize)),
				roleModel.WithKeywords(in.Keywords)).
			Order("id asc").Find(&items)

		for _, item := range items {
			var it sysPb.SysRole
			_ = copier.Copy(&it, item)

			list = append(list, &it)
		}
	}

	return &sysPb.RoleListResp{List: list, Total: total}, nil
}
