// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package sys

import (
	"context"

	"fast-boot/app/rpc/sys/sysPb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AnyReq              = sysPb.AnyReq
	DeptForm            = sysPb.DeptForm
	DeptListReq         = sysPb.DeptListReq
	DeptListResp        = sysPb.DeptListResp
	DeptOption          = sysPb.DeptOption
	DeptOptionResp      = sysPb.DeptOptionResp
	DictForm            = sysPb.DictForm
	DictListReq         = sysPb.DictListReq
	DictListResp        = sysPb.DictListResp
	DictOption          = sysPb.DictOption
	DictOptionsResp     = sysPb.DictOptionsResp
	DictTypeForm        = sysPb.DictTypeForm
	DictTypeListReq     = sysPb.DictTypeListReq
	DictTypeListResp    = sysPb.DictTypeListResp
	DictTypeOption      = sysPb.DictTypeOption
	DictTypeOptionsResp = sysPb.DictTypeOptionsResp
	IdReq               = sysPb.IdReq
	IdResp              = sysPb.IdResp
	IdsReq              = sysPb.IdsReq
	LoginReq            = sysPb.LoginReq
	LoginResp           = sysPb.LoginResp
	MenuForm            = sysPb.MenuForm
	MenuListReq         = sysPb.MenuListReq
	MenuListResp        = sysPb.MenuListResp
	MenuOption          = sysPb.MenuOption
	MenuOptionsResp     = sysPb.MenuOptionsResp
	RefreshTokenReq     = sysPb.RefreshTokenReq
	RoleForm            = sysPb.RoleForm
	RoleListReq         = sysPb.RoleListReq
	RoleListResp        = sysPb.RoleListResp
	RoleMenuIdsResp     = sysPb.RoleMenuIdsResp
	RoleOption          = sysPb.RoleOption
	RoleOptionsResp     = sysPb.RoleOptionsResp
	RoleSetMenuIdsReq   = sysPb.RoleSetMenuIdsReq
	RoutesReq           = sysPb.RoutesReq
	RoutesResp          = sysPb.RoutesResp
	SuccessIdResp       = sysPb.SuccessIdResp
	SuccessResp         = sysPb.SuccessResp
	SysDept             = sysPb.SysDept
	SysDeptItem         = sysPb.SysDeptItem
	SysDict             = sysPb.SysDict
	SysDictType         = sysPb.SysDictType
	SysMenu             = sysPb.SysMenu
	SysMenuItem         = sysPb.SysMenuItem
	SysRole             = sysPb.SysRole
	SysRoleMenu         = sysPb.SysRoleMenu
	SysUser             = sysPb.SysUser
	SysUserRole         = sysPb.SysUserRole
	TypeReq             = sysPb.TypeReq
	UserAddReq          = sysPb.UserAddReq
	UserChangePwdReq    = sysPb.UserChangePwdReq
	UserGetResp         = sysPb.UserGetResp
	UserListReq         = sysPb.UserListReq
	UserListResp        = sysPb.UserListResp
	UserUpdateReq       = sysPb.UserUpdateReq

	Sys interface {
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		RefreshToken(ctx context.Context, in *RefreshTokenReq, opts ...grpc.CallOption) (*LoginResp, error)
		UserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error)
		UserGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*UserGetResp, error)
		UserAdd(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*IdResp, error)
		UserUpdate(ctx context.Context, in *UserUpdateReq, opts ...grpc.CallOption) (*SuccessResp, error)
		UserChangePwd(ctx context.Context, in *UserChangePwdReq, opts ...grpc.CallOption) (*SuccessResp, error)
		UserDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error)
		RoleList(ctx context.Context, in *RoleListReq, opts ...grpc.CallOption) (*RoleListResp, error)
		RoleOptions(ctx context.Context, in *AnyReq, opts ...grpc.CallOption) (*RoleOptionsResp, error)
		RoleGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*SysRole, error)
		RoleAdd(ctx context.Context, in *RoleForm, opts ...grpc.CallOption) (*IdResp, error)
		RoleUpdate(ctx context.Context, in *RoleForm, opts ...grpc.CallOption) (*SuccessResp, error)
		RoleDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error)
		RoleMenuIds(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*RoleMenuIdsResp, error)
		RoleSetMenuIds(ctx context.Context, in *RoleSetMenuIdsReq, opts ...grpc.CallOption) (*SuccessResp, error)
		Routes(ctx context.Context, in *RoutesReq, opts ...grpc.CallOption) (*RoutesResp, error)
		MenuList(ctx context.Context, in *MenuListReq, opts ...grpc.CallOption) (*MenuListResp, error)
		MenuGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*SysMenu, error)
		MenuAdd(ctx context.Context, in *MenuForm, opts ...grpc.CallOption) (*SuccessResp, error)
		MenuUpdate(ctx context.Context, in *MenuForm, opts ...grpc.CallOption) (*SuccessResp, error)
		MenuDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error)
		MenuOptions(ctx context.Context, in *AnyReq, opts ...grpc.CallOption) (*MenuOptionsResp, error)
		DeptAdd(ctx context.Context, in *DeptForm, opts ...grpc.CallOption) (*SuccessResp, error)
		DeptUpdate(ctx context.Context, in *DeptForm, opts ...grpc.CallOption) (*SuccessResp, error)
		DeptDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error)
		DeptGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*SysDept, error)
		DeptList(ctx context.Context, in *DeptListReq, opts ...grpc.CallOption) (*DeptListResp, error)
		DeptOptions(ctx context.Context, in *AnyReq, opts ...grpc.CallOption) (*DeptOptionResp, error)
		DictTypeAdd(ctx context.Context, in *DictTypeForm, opts ...grpc.CallOption) (*SuccessResp, error)
		DictTypeUpdate(ctx context.Context, in *DictTypeForm, opts ...grpc.CallOption) (*SuccessResp, error)
		DictTypeDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error)
		DictTypeGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*SysDictType, error)
		DictTypeList(ctx context.Context, in *DictTypeListReq, opts ...grpc.CallOption) (*DictTypeListResp, error)
		DictTypeOptions(ctx context.Context, in *AnyReq, opts ...grpc.CallOption) (*DictTypeOptionsResp, error)
		DictAdd(ctx context.Context, in *DictForm, opts ...grpc.CallOption) (*SuccessResp, error)
		DictUpdate(ctx context.Context, in *DictForm, opts ...grpc.CallOption) (*SuccessResp, error)
		DictDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error)
		DictGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*SysDict, error)
		DictList(ctx context.Context, in *DictListReq, opts ...grpc.CallOption) (*DictListResp, error)
		DictOptions(ctx context.Context, in *TypeReq, opts ...grpc.CallOption) (*DictOptionsResp, error)
	}

	defaultSys struct {
		cli zrpc.Client
	}
)

func NewSys(cli zrpc.Client) Sys {
	return &defaultSys{
		cli: cli,
	}
}

func (m *defaultSys) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultSys) RefreshToken(ctx context.Context, in *RefreshTokenReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.RefreshToken(ctx, in, opts...)
}

func (m *defaultSys) UserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.UserList(ctx, in, opts...)
}

func (m *defaultSys) UserGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*UserGetResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.UserGet(ctx, in, opts...)
}

func (m *defaultSys) UserAdd(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*IdResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.UserAdd(ctx, in, opts...)
}

func (m *defaultSys) UserUpdate(ctx context.Context, in *UserUpdateReq, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.UserUpdate(ctx, in, opts...)
}

func (m *defaultSys) UserChangePwd(ctx context.Context, in *UserChangePwdReq, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.UserChangePwd(ctx, in, opts...)
}

func (m *defaultSys) UserDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.UserDel(ctx, in, opts...)
}

func (m *defaultSys) RoleList(ctx context.Context, in *RoleListReq, opts ...grpc.CallOption) (*RoleListResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.RoleList(ctx, in, opts...)
}

func (m *defaultSys) RoleOptions(ctx context.Context, in *AnyReq, opts ...grpc.CallOption) (*RoleOptionsResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.RoleOptions(ctx, in, opts...)
}

func (m *defaultSys) RoleGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*SysRole, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.RoleGet(ctx, in, opts...)
}

func (m *defaultSys) RoleAdd(ctx context.Context, in *RoleForm, opts ...grpc.CallOption) (*IdResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.RoleAdd(ctx, in, opts...)
}

func (m *defaultSys) RoleUpdate(ctx context.Context, in *RoleForm, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.RoleUpdate(ctx, in, opts...)
}

func (m *defaultSys) RoleDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.RoleDel(ctx, in, opts...)
}

func (m *defaultSys) RoleMenuIds(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*RoleMenuIdsResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.RoleMenuIds(ctx, in, opts...)
}

func (m *defaultSys) RoleSetMenuIds(ctx context.Context, in *RoleSetMenuIdsReq, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.RoleSetMenuIds(ctx, in, opts...)
}

func (m *defaultSys) Routes(ctx context.Context, in *RoutesReq, opts ...grpc.CallOption) (*RoutesResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.Routes(ctx, in, opts...)
}

func (m *defaultSys) MenuList(ctx context.Context, in *MenuListReq, opts ...grpc.CallOption) (*MenuListResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.MenuList(ctx, in, opts...)
}

func (m *defaultSys) MenuGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*SysMenu, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.MenuGet(ctx, in, opts...)
}

func (m *defaultSys) MenuAdd(ctx context.Context, in *MenuForm, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.MenuAdd(ctx, in, opts...)
}

func (m *defaultSys) MenuUpdate(ctx context.Context, in *MenuForm, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.MenuUpdate(ctx, in, opts...)
}

func (m *defaultSys) MenuDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.MenuDel(ctx, in, opts...)
}

func (m *defaultSys) MenuOptions(ctx context.Context, in *AnyReq, opts ...grpc.CallOption) (*MenuOptionsResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.MenuOptions(ctx, in, opts...)
}

func (m *defaultSys) DeptAdd(ctx context.Context, in *DeptForm, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.DeptAdd(ctx, in, opts...)
}

func (m *defaultSys) DeptUpdate(ctx context.Context, in *DeptForm, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.DeptUpdate(ctx, in, opts...)
}

func (m *defaultSys) DeptDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.DeptDel(ctx, in, opts...)
}

func (m *defaultSys) DeptGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*SysDept, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.DeptGet(ctx, in, opts...)
}

func (m *defaultSys) DeptList(ctx context.Context, in *DeptListReq, opts ...grpc.CallOption) (*DeptListResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.DeptList(ctx, in, opts...)
}

func (m *defaultSys) DeptOptions(ctx context.Context, in *AnyReq, opts ...grpc.CallOption) (*DeptOptionResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.DeptOptions(ctx, in, opts...)
}

func (m *defaultSys) DictTypeAdd(ctx context.Context, in *DictTypeForm, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.DictTypeAdd(ctx, in, opts...)
}

func (m *defaultSys) DictTypeUpdate(ctx context.Context, in *DictTypeForm, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.DictTypeUpdate(ctx, in, opts...)
}

func (m *defaultSys) DictTypeDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.DictTypeDel(ctx, in, opts...)
}

func (m *defaultSys) DictTypeGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*SysDictType, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.DictTypeGet(ctx, in, opts...)
}

func (m *defaultSys) DictTypeList(ctx context.Context, in *DictTypeListReq, opts ...grpc.CallOption) (*DictTypeListResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.DictTypeList(ctx, in, opts...)
}

func (m *defaultSys) DictTypeOptions(ctx context.Context, in *AnyReq, opts ...grpc.CallOption) (*DictTypeOptionsResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.DictTypeOptions(ctx, in, opts...)
}

func (m *defaultSys) DictAdd(ctx context.Context, in *DictForm, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.DictAdd(ctx, in, opts...)
}

func (m *defaultSys) DictUpdate(ctx context.Context, in *DictForm, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.DictUpdate(ctx, in, opts...)
}

func (m *defaultSys) DictDel(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*SuccessResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.DictDel(ctx, in, opts...)
}

func (m *defaultSys) DictGet(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*SysDict, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.DictGet(ctx, in, opts...)
}

func (m *defaultSys) DictList(ctx context.Context, in *DictListReq, opts ...grpc.CallOption) (*DictListResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.DictList(ctx, in, opts...)
}

func (m *defaultSys) DictOptions(ctx context.Context, in *TypeReq, opts ...grpc.CallOption) (*DictOptionsResp, error) {
	client := sysPb.NewSysClient(m.cli.Conn())
	return client.DictOptions(ctx, in, opts...)
}
