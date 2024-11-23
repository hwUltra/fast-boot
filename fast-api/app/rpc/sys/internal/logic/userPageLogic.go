package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"
	"fast-boot/common/static"
	"fmt"
	"github.com/hwUltra/fb-tools/gormx"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPageLogic {
	return &UserPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserPageLogic) UserPage(in *sysPb.UserPageReq) (*sysPb.UserPageResp, error) {

	userModel := model.SysUserModel{}
	total := int64(0)

	//适配Postgresql
	deptIdWhere := ""
	if l.svcCtx.Config.Gorm.SqlType == gormx.PostgresqlType {
		deptIdWhere = fmt.Sprint("concat(',',concat(sys_dept.tree_path::text,',',sys_dept.id::int),',') like concat('%,',?::int,',%')")
	} else {
		deptIdWhere = fmt.Sprint("concat(',',concat(sys_dept.tree_path,',',sys_dept.id),',') like concat('%,',?,',%')")
	}

	if err := l.svcCtx.GormClient.GormDb.Model(userModel).
		Joins("left join sys_dept on sys_user.dept_id = sys_dept.id").
		Scopes(
			userModel.WithStatus(in.Status),
			userModel.WithCreatedAt(in.StartTime, in.EndTime),
			userModel.WithKeywords(in.Keywords)).
		Where(deptIdWhere, in.DeptId).
		Count(&total).Error; err != nil {
		return nil, err
	}

	list := make([]*sysPb.SysUser, 0)

	if total > 0 {
		items := make([]*model.SysUserModel, 0)
		l.svcCtx.GormClient.GormDb.Model(userModel).
			Joins("left join sys_dept on sys_user.dept_id = sys_dept.id").
			Preload("Dept", "status = 1").
			Preload("Roles", "status = 1").
			Scopes(
				gormx.Paginate(int(in.PageNum), int(in.PageSize)),
				userModel.WithStatus(in.Status),
				userModel.WithCreatedAt(in.StartTime, in.EndTime),
				userModel.WithKeywords(in.Keywords)).
			Where(deptIdWhere, in.DeptId).
			Order("id asc").Find(&items)

		for _, item := range items {
			var it sysPb.SysUser
			_ = copier.Copy(&it, item)
			if item.Dept != nil {
				it.DeptName = item.Dept.Name
			}
			it.GenderLabel = static.GetGenderStr(it.Gender)

			roles := make([]string, 0)
			for _, v := range item.Roles {
				roles = append(roles, v.Code)
			}
			it.Roles = roles

			list = append(list, &it)
		}
	}

	return &sysPb.UserPageResp{List: list, Total: total}, nil
}
