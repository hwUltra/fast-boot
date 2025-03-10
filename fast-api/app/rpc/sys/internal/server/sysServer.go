// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: sys.proto

package server

import (
	"context"

	"fast-boot/app/rpc/sys/internal/logic"
	"fast-boot/app/rpc/sys/internal/svc"
	"fast-boot/app/rpc/sys/sysPb"
)

type SysServer struct {
	svcCtx *svc.ServiceContext
	sysPb.UnimplementedSysServer
}

func NewSysServer(svcCtx *svc.ServiceContext) *SysServer {
	return &SysServer{
		svcCtx: svcCtx,
	}
}

func (s *SysServer) Login(ctx context.Context, in *sysPb.LoginReq) (*sysPb.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *SysServer) RefreshToken(ctx context.Context, in *sysPb.RefreshTokenReq) (*sysPb.LoginResp, error) {
	l := logic.NewRefreshTokenLogic(ctx, s.svcCtx)
	return l.RefreshToken(in)
}

func (s *SysServer) UserPage(ctx context.Context, in *sysPb.UserPageReq) (*sysPb.UserPageResp, error) {
	l := logic.NewUserPageLogic(ctx, s.svcCtx)
	return l.UserPage(in)
}

func (s *SysServer) UserGet(ctx context.Context, in *sysPb.IdReq) (*sysPb.UserGetResp, error) {
	l := logic.NewUserGetLogic(ctx, s.svcCtx)
	return l.UserGet(in)
}

func (s *SysServer) UserProfile(ctx context.Context, in *sysPb.IdReq) (*sysPb.SysUserProfile, error) {
	l := logic.NewUserProfileLogic(ctx, s.svcCtx)
	return l.UserProfile(in)
}

func (s *SysServer) UserAdd(ctx context.Context, in *sysPb.UserAddReq) (*sysPb.IdResp, error) {
	l := logic.NewUserAddLogic(ctx, s.svcCtx)
	return l.UserAdd(in)
}

func (s *SysServer) UserUpdate(ctx context.Context, in *sysPb.UserUpdateReq) (*sysPb.SuccessResp, error) {
	l := logic.NewUserUpdateLogic(ctx, s.svcCtx)
	return l.UserUpdate(in)
}

func (s *SysServer) UserChangePwd(ctx context.Context, in *sysPb.UserChangePwdReq) (*sysPb.SuccessResp, error) {
	l := logic.NewUserChangePwdLogic(ctx, s.svcCtx)
	return l.UserChangePwd(in)
}

func (s *SysServer) UserDel(ctx context.Context, in *sysPb.IdsReq) (*sysPb.SuccessResp, error) {
	l := logic.NewUserDelLogic(ctx, s.svcCtx)
	return l.UserDel(in)
}

func (s *SysServer) RoleList(ctx context.Context, in *sysPb.RoleListReq) (*sysPb.RoleListResp, error) {
	l := logic.NewRoleListLogic(ctx, s.svcCtx)
	return l.RoleList(in)
}

func (s *SysServer) RoleOptions(ctx context.Context, in *sysPb.AnyReq) (*sysPb.RoleOptionsResp, error) {
	l := logic.NewRoleOptionsLogic(ctx, s.svcCtx)
	return l.RoleOptions(in)
}

func (s *SysServer) RoleGet(ctx context.Context, in *sysPb.IdReq) (*sysPb.SysRole, error) {
	l := logic.NewRoleGetLogic(ctx, s.svcCtx)
	return l.RoleGet(in)
}

func (s *SysServer) RoleAdd(ctx context.Context, in *sysPb.RoleForm) (*sysPb.IdResp, error) {
	l := logic.NewRoleAddLogic(ctx, s.svcCtx)
	return l.RoleAdd(in)
}

func (s *SysServer) RoleUpdate(ctx context.Context, in *sysPb.RoleForm) (*sysPb.SuccessResp, error) {
	l := logic.NewRoleUpdateLogic(ctx, s.svcCtx)
	return l.RoleUpdate(in)
}

func (s *SysServer) RoleDel(ctx context.Context, in *sysPb.IdsReq) (*sysPb.SuccessResp, error) {
	l := logic.NewRoleDelLogic(ctx, s.svcCtx)
	return l.RoleDel(in)
}

func (s *SysServer) RoleMenuIds(ctx context.Context, in *sysPb.IdReq) (*sysPb.RoleMenuIdsResp, error) {
	l := logic.NewRoleMenuIdsLogic(ctx, s.svcCtx)
	return l.RoleMenuIds(in)
}

func (s *SysServer) RoleSetMenuIds(ctx context.Context, in *sysPb.RoleSetMenuIdsReq) (*sysPb.SuccessResp, error) {
	l := logic.NewRoleSetMenuIdsLogic(ctx, s.svcCtx)
	return l.RoleSetMenuIds(in)
}

func (s *SysServer) Routes(ctx context.Context, in *sysPb.RoutesReq) (*sysPb.RoutesResp, error) {
	l := logic.NewRoutesLogic(ctx, s.svcCtx)
	return l.Routes(in)
}

func (s *SysServer) MenuList(ctx context.Context, in *sysPb.MenuListReq) (*sysPb.MenuListResp, error) {
	l := logic.NewMenuListLogic(ctx, s.svcCtx)
	return l.MenuList(in)
}

func (s *SysServer) MenuGet(ctx context.Context, in *sysPb.IdReq) (*sysPb.SysMenu, error) {
	l := logic.NewMenuGetLogic(ctx, s.svcCtx)
	return l.MenuGet(in)
}

func (s *SysServer) MenuAdd(ctx context.Context, in *sysPb.MenuForm) (*sysPb.SuccessResp, error) {
	l := logic.NewMenuAddLogic(ctx, s.svcCtx)
	return l.MenuAdd(in)
}

func (s *SysServer) MenuUpdate(ctx context.Context, in *sysPb.MenuForm) (*sysPb.SuccessResp, error) {
	l := logic.NewMenuUpdateLogic(ctx, s.svcCtx)
	return l.MenuUpdate(in)
}

func (s *SysServer) MenuDel(ctx context.Context, in *sysPb.IdsReq) (*sysPb.SuccessResp, error) {
	l := logic.NewMenuDelLogic(ctx, s.svcCtx)
	return l.MenuDel(in)
}

func (s *SysServer) MenuOptions(ctx context.Context, in *sysPb.AnyReq) (*sysPb.MenuOptionsResp, error) {
	l := logic.NewMenuOptionsLogic(ctx, s.svcCtx)
	return l.MenuOptions(in)
}

func (s *SysServer) DeptAdd(ctx context.Context, in *sysPb.DeptForm) (*sysPb.SuccessResp, error) {
	l := logic.NewDeptAddLogic(ctx, s.svcCtx)
	return l.DeptAdd(in)
}

func (s *SysServer) DeptUpdate(ctx context.Context, in *sysPb.DeptForm) (*sysPb.SuccessResp, error) {
	l := logic.NewDeptUpdateLogic(ctx, s.svcCtx)
	return l.DeptUpdate(in)
}

func (s *SysServer) DeptDel(ctx context.Context, in *sysPb.IdsReq) (*sysPb.SuccessResp, error) {
	l := logic.NewDeptDelLogic(ctx, s.svcCtx)
	return l.DeptDel(in)
}

func (s *SysServer) DeptGet(ctx context.Context, in *sysPb.IdReq) (*sysPb.SysDept, error) {
	l := logic.NewDeptGetLogic(ctx, s.svcCtx)
	return l.DeptGet(in)
}

func (s *SysServer) DeptList(ctx context.Context, in *sysPb.DeptListReq) (*sysPb.DeptListResp, error) {
	l := logic.NewDeptListLogic(ctx, s.svcCtx)
	return l.DeptList(in)
}

func (s *SysServer) DeptOptions(ctx context.Context, in *sysPb.AnyReq) (*sysPb.DeptOptionResp, error) {
	l := logic.NewDeptOptionsLogic(ctx, s.svcCtx)
	return l.DeptOptions(in)
}

func (s *SysServer) DictSimList(ctx context.Context, in *sysPb.AnyReq) (*sysPb.DictSimListResp, error) {
	l := logic.NewDictSimListLogic(ctx, s.svcCtx)
	return l.DictSimList(in)
}

func (s *SysServer) DictAdd(ctx context.Context, in *sysPb.DictForm) (*sysPb.SuccessResp, error) {
	l := logic.NewDictAddLogic(ctx, s.svcCtx)
	return l.DictAdd(in)
}

func (s *SysServer) DictUpdate(ctx context.Context, in *sysPb.DictForm) (*sysPb.SuccessResp, error) {
	l := logic.NewDictUpdateLogic(ctx, s.svcCtx)
	return l.DictUpdate(in)
}

func (s *SysServer) DictDel(ctx context.Context, in *sysPb.IdsReq) (*sysPb.SuccessResp, error) {
	l := logic.NewDictDelLogic(ctx, s.svcCtx)
	return l.DictDel(in)
}

func (s *SysServer) DictGet(ctx context.Context, in *sysPb.IdReq) (*sysPb.SysDict, error) {
	l := logic.NewDictGetLogic(ctx, s.svcCtx)
	return l.DictGet(in)
}

func (s *SysServer) DictPage(ctx context.Context, in *sysPb.DictPageReq) (*sysPb.DictPageResp, error) {
	l := logic.NewDictPageLogic(ctx, s.svcCtx)
	return l.DictPage(in)
}

func (s *SysServer) DictDataAdd(ctx context.Context, in *sysPb.DictDataForm) (*sysPb.SuccessResp, error) {
	l := logic.NewDictDataAddLogic(ctx, s.svcCtx)
	return l.DictDataAdd(in)
}

func (s *SysServer) DictDataUpdate(ctx context.Context, in *sysPb.DictDataForm) (*sysPb.SuccessResp, error) {
	l := logic.NewDictDataUpdateLogic(ctx, s.svcCtx)
	return l.DictDataUpdate(in)
}

func (s *SysServer) DictDataDel(ctx context.Context, in *sysPb.IdsReq) (*sysPb.SuccessResp, error) {
	l := logic.NewDictDataDelLogic(ctx, s.svcCtx)
	return l.DictDataDel(in)
}

func (s *SysServer) DictDataGet(ctx context.Context, in *sysPb.IdReq) (*sysPb.SysDictData, error) {
	l := logic.NewDictDataGetLogic(ctx, s.svcCtx)
	return l.DictDataGet(in)
}

func (s *SysServer) DictDataPage(ctx context.Context, in *sysPb.DictDataPageReq) (*sysPb.DictDataPageResp, error) {
	l := logic.NewDictDataPageLogic(ctx, s.svcCtx)
	return l.DictDataPage(in)
}

func (s *SysServer) SysNoticePage(ctx context.Context, in *sysPb.SysNoticePageReq) (*sysPb.SysNoticePageResp, error) {
	l := logic.NewSysNoticePageLogic(ctx, s.svcCtx)
	return l.SysNoticePage(in)
}
