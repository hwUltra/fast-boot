// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	auth "fast-boot/app/api/admin/internal/handler/auth"
	omsorder "fast-boot/app/api/admin/internal/handler/oms/order"
	other "fast-boot/app/api/admin/internal/handler/other"
	pmsgoods "fast-boot/app/api/admin/internal/handler/pms/goods"
	pmsshop "fast-boot/app/api/admin/internal/handler/pms/shop"
	sysdept "fast-boot/app/api/admin/internal/handler/sys/dept"
	sysdict "fast-boot/app/api/admin/internal/handler/sys/dict"
	syslogs "fast-boot/app/api/admin/internal/handler/sys/logs"
	sysmenu "fast-boot/app/api/admin/internal/handler/sys/menu"
	sysnotice "fast-boot/app/api/admin/internal/handler/sys/notice"
	sysrole "fast-boot/app/api/admin/internal/handler/sys/role"
	sysuser "fast-boot/app/api/admin/internal/handler/sys/user"
	umsuser "fast-boot/app/api/admin/internal/handler/ums/user"
	"fast-boot/app/api/admin/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 登录
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: auth.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/auth"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 退出登录
				Method:  http.MethodGet,
				Path:    "/logout",
				Handler: auth.LogoutHandler(serverCtx),
			},
			{
				// 获取个人信息
				Method:  http.MethodGet,
				Path:    "/me",
				Handler: auth.MeHandler(serverCtx),
			},
			{
				// 刷新ToKen
				Method:  http.MethodGet,
				Path:    "/refreshToken",
				Handler: auth.RefreshTokenHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/auth"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 创建订单
				Method:  http.MethodPost,
				Path:    "/",
				Handler: omsorder.CreateHandler(serverCtx),
			},
			{
				// 订单列表
				Method:  http.MethodGet,
				Path:    "/",
				Handler: omsorder.PageHandler(serverCtx),
			},
			{
				// 删除订单
				Method:  http.MethodDelete,
				Path:    "/:id",
				Handler: omsorder.DelHandler(serverCtx),
			},
			{
				// 修改订单
				Method:  http.MethodPut,
				Path:    "/:id",
				Handler: omsorder.UpdateHandler(serverCtx),
			},
			{
				// 获取订单
				Method:  http.MethodGet,
				Path:    "/:id",
				Handler: omsorder.GetHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/order"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 验证码
				Method:  http.MethodGet,
				Path:    "/captcha",
				Handler: other.CaptchaHandler(serverCtx),
			},
			{
				// 文件上传
				Method:  http.MethodPost,
				Path:    "/upload",
				Handler: other.UploadHandler(serverCtx),
			},
		},
		rest.WithPrefix("/other"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 商品列表
				Method:  http.MethodGet,
				Path:    "/",
				Handler: pmsgoods.PageHandler(serverCtx),
			},
			{
				// 获取商品
				Method:  http.MethodGet,
				Path:    "/:id",
				Handler: pmsgoods.GetHandler(serverCtx),
			},
			{
				// 删除
				Method:  http.MethodDelete,
				Path:    "/:ids",
				Handler: pmsgoods.DelHandler(serverCtx),
			},
			{
				// 新增商品
				Method:  http.MethodPost,
				Path:    "/edit",
				Handler: pmsgoods.EditHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/goods"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 新增类型
				Method:  http.MethodPost,
				Path:    "/",
				Handler: pmsshop.AddHandler(serverCtx),
			},
			{
				// 店铺列表
				Method:  http.MethodGet,
				Path:    "/",
				Handler: pmsshop.PageHandler(serverCtx),
			},
			{
				// 修改类型
				Method:  http.MethodPut,
				Path:    "/:id",
				Handler: pmsshop.UpdateHandler(serverCtx),
			},
			{
				// 获取店铺
				Method:  http.MethodGet,
				Path:    "/:id",
				Handler: pmsshop.GetHandler(serverCtx),
			},
			{
				// 删除
				Method:  http.MethodDelete,
				Path:    "/:ids",
				Handler: pmsshop.DelHandler(serverCtx),
			},
			{
				// 新增Brand
				Method:  http.MethodPost,
				Path:    "/brand",
				Handler: pmsshop.BrandAddHandler(serverCtx),
			},
			{
				// 修改Brand
				Method:  http.MethodPut,
				Path:    "/brand",
				Handler: pmsshop.BrandUpdateHandler(serverCtx),
			},
			{
				// Brand列表
				Method:  http.MethodGet,
				Path:    "/brand",
				Handler: pmsshop.BrandPageHandler(serverCtx),
			},
			{
				// 获取Brand
				Method:  http.MethodGet,
				Path:    "/brand/:id",
				Handler: pmsshop.BrandGetHandler(serverCtx),
			},
			{
				// 删除
				Method:  http.MethodDelete,
				Path:    "/brand/:ids",
				Handler: pmsshop.BrandDelHandler(serverCtx),
			},
			{
				// Brand下拉列表
				Method:  http.MethodGet,
				Path:    "/brand/options/:id",
				Handler: pmsshop.BrandOptionsHandler(serverCtx),
			},
			{
				// 新增类型
				Method:  http.MethodPost,
				Path:    "/category",
				Handler: pmsshop.CategoryAddHandler(serverCtx),
			},
			{
				// 类型列表
				Method:  http.MethodGet,
				Path:    "/category",
				Handler: pmsshop.CategoryListHandler(serverCtx),
			},
			{
				// 修改类型
				Method:  http.MethodPut,
				Path:    "/category/:id",
				Handler: pmsshop.CategoryUpdateHandler(serverCtx),
			},
			{
				// 获取店铺
				Method:  http.MethodGet,
				Path:    "/category/:id",
				Handler: pmsshop.CategoryGetHandler(serverCtx),
			},
			{
				// 删除类型
				Method:  http.MethodDelete,
				Path:    "/category/:ids",
				Handler: pmsshop.CategoryDelHandler(serverCtx),
			},
			{
				// 新增类型-属性
				Method:  http.MethodPost,
				Path:    "/category/attribute",
				Handler: pmsshop.CategoryAttributeAddHandler(serverCtx),
			},
			{
				// 新增类型-属性
				Method:  http.MethodGet,
				Path:    "/category/attribute",
				Handler: pmsshop.CategoryAttributeListHandler(serverCtx),
			},
			{
				// 类型下拉列表
				Method:  http.MethodGet,
				Path:    "/category/options/:id",
				Handler: pmsshop.CategoryOptionsHandler(serverCtx),
			},
			{
				// 店铺下拉列表
				Method:  http.MethodGet,
				Path:    "/options",
				Handler: pmsshop.OptionsHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/shop"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 新增部门
				Method:  http.MethodPost,
				Path:    "/",
				Handler: sysdept.AddHandler(serverCtx),
			},
			{
				// 部门分页列表
				Method:  http.MethodGet,
				Path:    "/",
				Handler: sysdept.ListHandler(serverCtx),
			},
			{
				// 修改部门
				Method:  http.MethodPut,
				Path:    "/:id",
				Handler: sysdept.UpdateHandler(serverCtx),
			},
			{
				// 获取部门
				Method:  http.MethodGet,
				Path:    "/:id",
				Handler: sysdept.GetHandler(serverCtx),
			},
			{
				// 删除部门
				Method:  http.MethodDelete,
				Path:    "/:ids",
				Handler: sysdept.DelHandler(serverCtx),
			},
			{
				// 部门下拉列表
				Method:  http.MethodGet,
				Path:    "/options",
				Handler: sysdept.OptionsHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/sys/dept"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: sysdict.AddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: sysdict.PageHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/:id",
				Handler: sysdict.UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/:id",
				Handler: sysdict.GetHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/:ids",
				Handler: sysdict.DelHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/data",
				Handler: sysdict.DataAddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/data",
				Handler: sysdict.DataPageHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/data/:id",
				Handler: sysdict.DataUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/data/:id",
				Handler: sysdict.DataGetHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/data/:ids",
				Handler: sysdict.DataDelHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/simlist",
				Handler: sysdict.SimListHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/sys/dict"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: syslogs.AddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: syslogs.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/visit-stats",
				Handler: syslogs.VisitStatsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/visit-trend",
				Handler: syslogs.VisitTrendHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/sys/logs"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 菜单列表
				Method:  http.MethodGet,
				Path:    "/",
				Handler: sysmenu.PageHandler(serverCtx),
			},
			{
				// 新增菜单
				Method:  http.MethodPost,
				Path:    "/",
				Handler: sysmenu.AddHandler(serverCtx),
			},
			{
				// 修改菜单
				Method:  http.MethodPut,
				Path:    "/:id",
				Handler: sysmenu.UpdateHandler(serverCtx),
			},
			{
				// 获取菜单
				Method:  http.MethodGet,
				Path:    "/:id",
				Handler: sysmenu.GetHandler(serverCtx),
			},
			{
				// 删除角色
				Method:  http.MethodDelete,
				Path:    "/:ids",
				Handler: sysmenu.DelHandler(serverCtx),
			},
			{
				// 菜单下拉列表
				Method:  http.MethodGet,
				Path:    "/options",
				Handler: sysmenu.OptionsHandler(serverCtx),
			},
			{
				// 路由列表
				Method:  http.MethodGet,
				Path:    "/routes",
				Handler: sysmenu.RoutesHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/sys/menu"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: sysnotice.AddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: sysnotice.PageHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/my-page",
				Handler: sysnotice.MyPageHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/sys/notice"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 角色分页列表
				Method:  http.MethodGet,
				Path:    "/",
				Handler: sysrole.PageHandler(serverCtx),
			},
			{
				// 新增角色
				Method:  http.MethodPost,
				Path:    "/",
				Handler: sysrole.AddHandler(serverCtx),
			},
			{
				// 修改角色
				Method:  http.MethodPut,
				Path:    "/:id",
				Handler: sysrole.UpdateHandler(serverCtx),
			},
			{
				// 获取角色
				Method:  http.MethodGet,
				Path:    "/:id",
				Handler: sysrole.GetHandler(serverCtx),
			},
			{
				// 删除角色
				Method:  http.MethodDelete,
				Path:    "/:ids",
				Handler: sysrole.DelHandler(serverCtx),
			},
			{
				// 修改用户状态
				Method:  http.MethodPost,
				Path:    "/changeStatus",
				Handler: sysrole.ChangeStatusHandler(serverCtx),
			},
			{
				// 获取角色的菜单ID集合
				Method:  http.MethodGet,
				Path:    "/menuIds/:id",
				Handler: sysrole.MenuIdsHandler(serverCtx),
			},
			{
				// 角色下拉列表
				Method:  http.MethodGet,
				Path:    "/options",
				Handler: sysrole.OptionsHandler(serverCtx),
			},
			{
				// 分配菜单权限给角色
				Method:  http.MethodPost,
				Path:    "/setMenuIds",
				Handler: sysrole.SetMenuIdsHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/sys/role"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 新增用户
				Method:  http.MethodPost,
				Path:    "/",
				Handler: sysuser.AddHandler(serverCtx),
			},
			{
				// 用户分页列表
				Method:  http.MethodGet,
				Path:    "/",
				Handler: sysuser.PageHandler(serverCtx),
			},
			{
				// 修改用户
				Method:  http.MethodPut,
				Path:    "/:id",
				Handler: sysuser.UpdateHandler(serverCtx),
			},
			{
				// 获取用户
				Method:  http.MethodGet,
				Path:    "/:id",
				Handler: sysuser.GetHandler(serverCtx),
			},
			{
				// 删除多个用户
				Method:  http.MethodDelete,
				Path:    "/:ids",
				Handler: sysuser.DelHandler(serverCtx),
			},
			{
				// 修改用户密码
				Method:  http.MethodPost,
				Path:    "/changePwd",
				Handler: sysuser.ChangePwdHandler(serverCtx),
			},
			{
				// 修改用户状态
				Method:  http.MethodPost,
				Path:    "/changeStatus",
				Handler: sysuser.ChangeStatusHandler(serverCtx),
			},
			{
				// Profile 获取用户
				Method:  http.MethodGet,
				Path:    "/profile",
				Handler: sysuser.ProfileHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/sys/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 添加用户
				Method:  http.MethodPost,
				Path:    "/",
				Handler: umsuser.AddHandler(serverCtx),
			},
			{
				// 获取用户列表
				Method:  http.MethodGet,
				Path:    "/",
				Handler: umsuser.PageHandler(serverCtx),
			},
			{
				// 修改用户信息
				Method:  http.MethodPut,
				Path:    "/:id",
				Handler: umsuser.UpdateHandler(serverCtx),
			},
			{
				// 获取用户信息
				Method:  http.MethodGet,
				Path:    "/:id",
				Handler: umsuser.GetHandler(serverCtx),
			},
			{
				// 删除用户
				Method:  http.MethodDelete,
				Path:    "/:ids",
				Handler: umsuser.DelHandler(serverCtx),
			},
			{
				// 获取第三方用户列表
				Method:  http.MethodGet,
				Path:    "/third",
				Handler: umsuser.ThirdPageHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/user"),
	)
}
