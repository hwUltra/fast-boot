package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"gorm.io/gorm"

	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleSetMenuIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleSetMenuIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleSetMenuIdsLogic {
	return &RoleSetMenuIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleSetMenuIdsLogic) RoleSetMenuIds(in *sysPb.RoleSetMenuIdsReq) (*sysPb.SuccessResp, error) {

	if err := l.svcCtx.GormConn.Transaction(func(tx *gorm.DB) error {
		//删除
		if err := tx.Where("role_id = ?", in.RoleId).Delete(&model.SysRoleMenuModel{}).Error; err != nil {
			return err
		}
		for _, menuId := range in.MenuIds {
			if err := tx.FirstOrCreate(&model.SysRoleMenuModel{}, model.SysRoleMenuModel{
				RoleID: in.RoleId,
				MenuID: menuId,
			}).Error; err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &sysPb.SuccessResp{}, nil
}
