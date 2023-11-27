package logic

import (
	"context"
	"fast-boot/app/rpc/model"
	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"
	"fast-boot/common/gormV2"
	"fast-boot/common/static"
	"fast-boot/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserList  列表
func (l *UserListLogic) UserList(in *sysPb.UserListReq) (*sysPb.UserListResp, error) {

	userModel := model.SysUserModel{}
	//total
	total := int64(0)
	//l.svcCtx.GormConn.Model(userModel).
	//	Scopes(userModel.WithStatus(in.Status),
	//		userModel.WithCreatedAt(in.StartTime, in.EndTime),
	//		userModel.WithKeywords(in.Keywords)).
	//	Joins("left join sys_dept on sys_user.dept_id = sys_dept.id").
	//	Where("concat(',',concat(d.tree_path,',',d.id),',') like concat('%,',?,',%')", in.DeptId)
	if err := l.svcCtx.GormConn.Model(userModel).
		Joins("left join `sys_dept` as d on `sys_user`.dept_id = d.id").
		Scopes(
			userModel.WithStatus(in.Status),
			userModel.WithCreatedAt(in.StartTime, in.EndTime),
			userModel.WithKeywords(in.Keywords)).
		Where("concat(',',concat(d.tree_path,',',d.id),',') like concat('%,',?,',%')", in.DeptId).
		Count(&total).Error; err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to get  err : %v , in :%+v", err, in)
	}

	list := make([]*sysPb.SysUser, 0)

	if total > 0 {
		items := make([]*model.SysUserModel, 0)
		l.svcCtx.GormConn.Model(userModel).
			Joins("left join `sys_dept` as d on `sys_user`.dept_id = d.id").
			Preload("Dept", "status = 1").
			Preload("Roles", "status = 1").
			Scopes(
				gormV2.Paginate(int(in.PageNum), int(in.PageSize)),
				userModel.WithStatus(in.Status),
				userModel.WithCreatedAt(in.StartTime, in.EndTime),
				userModel.WithKeywords(in.Keywords)).
			Where("concat(',',concat(d.tree_path,',',d.id),',') like concat('%,',?,',%')", in.DeptId).
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

	return &sysPb.UserListResp{List: list, Total: total}, nil
}
