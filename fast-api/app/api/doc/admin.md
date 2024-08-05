### 1. "登录"

1. route definition

- Url: /auth/login
- Method: POST
- Request: `LoginReq`
- Response: `TokenResp`

2. request definition



```golang
type LoginReq struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
```


3. response definition



```golang
type TokenResp struct {
	AccessToken string `json:"accessToken"`
	AccessExpire int64 `json:"accessExpire"`
	RefreshAfter int64 `json:"refreshAfter"`
}
```

### 2. "刷新ToKen"

1. route definition

- Url: /auth/refreshToken
- Method: GET
- Request: `-`
- Response: `TokenResp`

2. request definition



3. response definition



```golang
type TokenResp struct {
	AccessToken string `json:"accessToken"`
	AccessExpire int64 `json:"accessExpire"`
	RefreshAfter int64 `json:"refreshAfter"`
}
```

### 3. "获取个人信息"

1. route definition

- Url: /auth/me
- Method: GET
- Request: `-`
- Response: `MeResp`

2. request definition



3. response definition



```golang
type MeResp struct {
	UserId int64 `json:"userId"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar string `json:"avatar"`
	Roles []string `json:"roles,omitempty"`
	Perms []string `json:"perms,omitempty"`
}
```

### 4. "退出登录"

1. route definition

- Url: /auth/logout
- Method: GET
- Request: `-`
- Response: `-`

2. request definition



3. response definition


### 5. "用户分页列表"

1. route definition

- Url: /sys/user/list
- Method: GET
- Request: `SysUserListReq`
- Response: `SysUserListResp`

2. request definition



```golang
type SysUserListReq struct {
	Keywords string `form:"keywords,omitempty,optional"`
	DeptId int64 `form:"deptId,omitempty,optional"`
	Status int64 `form:"status,omitempty,optional,default=-1"`
	PageNum int64 `form:"pageNum,omitempty,optional,default=1"`
	PageSize int64 `form:"pageSize,omitempty,optional,default=10"`
	StartTime string `form:"startTime,optional"` // 开始时间
	EndTime string `form:"endTime,optional"` // 结束时间
}
```


3. response definition



```golang
type SysUserListResp struct {
	List []*SysUserInfo `json:"list"`
	Total int64 `json:"total"`
}
```

### 6. "新增用户"

1. route definition

- Url: /sys/user/add
- Method: POST
- Request: `SysUserFormReq`
- Response: `IdResp`

2. request definition



```golang
type SysUserFormReq struct {
	Id int64 `json:"id,optional"`
	Username string `json:"username"`
	Nickname string `json:"nickname,optional"`
	Gender int64 `json:"gender,optional"`
	Password string `json:"password,optional"`
	DeptId int64 `json:"deptId,optional"`
	Avatar string `json:"avatar,optional"`
	Mobile string `json:"mobile,optional"`
	Status int64 `json:"status,default=1"`
	Email string `json:"email,optional"`
	RoleIds []int64 `json:"roleIds"`
}
```


3. response definition



```golang
type IdResp struct {
	ID int64 `json:"id"`
}
```

### 7. "获取用户"

1. route definition

- Url: /sys/user/:id
- Method: GET
- Request: `PathIdReq`
- Response: `SysUserGet`

2. request definition



```golang
type PathIdReq struct {
	Id int64 `path:"id"`
}
```


3. response definition



```golang
type SysUserGet struct {
	Id int64 `json:"id"`
	Username string `json:"username"` // 用户名
	Nickname string `json:"nickname"` // 昵称
	Gender int64 `json:"gender"` // 性别((1:男;2:女))
	DeptId int64 `json:"deptId"` // 部门ID
	Avatar string `json:"avatar"` // 用户头像
	Mobile string `json:"mobile"` // 联系方式
	Status int64 `json:"status"` // 用户状态((1:正常;0:禁用))
	Email string `json:"email"` // 用户邮箱
	RoleIds []int64 `json:"roleIds"`
}
```

### 8. "修改用户"

1. route definition

- Url: /sys/user/update
- Method: POST
- Request: `SysUserFormReq`
- Response: `-`

2. request definition



```golang
type SysUserFormReq struct {
	Id int64 `json:"id,optional"`
	Username string `json:"username"`
	Nickname string `json:"nickname,optional"`
	Gender int64 `json:"gender,optional"`
	Password string `json:"password,optional"`
	DeptId int64 `json:"deptId,optional"`
	Avatar string `json:"avatar,optional"`
	Mobile string `json:"mobile,optional"`
	Status int64 `json:"status,default=1"`
	Email string `json:"email,optional"`
	RoleIds []int64 `json:"roleIds"`
}
```


3. response definition


### 9. "删除多个用户"

1. route definition

- Url: /sys/user/del/:ids
- Method: GET
- Request: `PathIdsReq`
- Response: `-`

2. request definition



```golang
type PathIdsReq struct {
	Ids string `path:"ids"`
}
```


3. response definition


### 10. "修改用户密码"

1. route definition

- Url: /sys/user/changePwd
- Method: POST
- Request: `SysUserChangePwdReq`
- Response: `-`

2. request definition



```golang
type SysUserChangePwdReq struct {
	UserId int64 `json:"userId"`
	Password string `json:"password"`
}
```


3. response definition


### 11. "修改用户状态"

1. route definition

- Url: /sys/user/changeStatus
- Method: POST
- Request: `SysUserChangeStatusReq`
- Response: `-`

2. request definition



```golang
type SysUserChangeStatusReq struct {
	UserId int64 `json:"userId"`
	Status string `json:"status"`
}
```


3. response definition


### 12. "路由列表"

1. route definition

- Url: /sys/menu/routes
- Method: GET
- Request: `-`
- Response: `RoutesResp`

2. request definition



3. response definition



```golang
type RoutesResp struct {
	List []*RouteData `json:"list"`
}
```

### 13. "菜单列表"

1. route definition

- Url: /sys/menu/list
- Method: GET
- Request: `SysMenuListReq`
- Response: `SysMenuListResp`

2. request definition



```golang
type SysMenuListReq struct {
	Keywords string `json:"keywords,omitempty,optional"`
}
```


3. response definition



```golang
type SysMenuListResp struct {
	List []*SysMenuItem `json:"list"`
}
```

### 14. "新增菜单"

1. route definition

- Url: /sys/menu/add
- Method: POST
- Request: `SysMenuForm`
- Response: `IdResp`

2. request definition



```golang
type SysMenuForm struct {
	Id int64 `json:"id,optional"`
	ParentId int64 `json:"parentId"`
	Name string `json:"name"`
	Type int64 `json:"type,optional"`
	Path string `json:"path,optional"`
	Component string `json:"component,optional"`
	Perm string `json:"perm,optional"`
	Visible int64 `json:"visible,optional"`
	Sort int64 `json:"sort,default=false"`
	Icon string `json:"icon,optional"`
	Redirect string `json:"redirect,optional"`
}
```


3. response definition



```golang
type IdResp struct {
	ID int64 `json:"id"`
}
```

### 15. "获取菜单"

1. route definition

- Url: /sys/menu/:id
- Method: GET
- Request: `PathIdReq`
- Response: `SysMenuInfo`

2. request definition



```golang
type PathIdReq struct {
	Id int64 `path:"id"`
}
```


3. response definition



```golang
type SysMenuInfo struct {
	Id int64 `json:"id"`
	ParentId int64 `json:"parentId"`
	TreePath string `json:"treePath"`
	Name string `json:"name"`
	Type int64 `json:"type"`
	Path string `json:"path"`
	Component string `json:"component"`
	Perm string `json:"perm"`
	Visible int64 `json:"visible"`
	Sort int64 `json:"sort,default=1"`
	Icon string `json:"icon"`
	Redirect string `json:"redirect"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
```

### 16. "菜单下拉列表"

1. route definition

- Url: /sys/menu/options
- Method: GET
- Request: `-`
- Response: `SysMenuOptionsResp`

2. request definition



3. response definition



```golang
type SysMenuOptionsResp struct {
	List []*SysMenuOption `json:"list"`
}
```

### 17. "修改菜单"

1. route definition

- Url: /sys/menu/update
- Method: POST
- Request: `SysMenuForm`
- Response: `IdResp`

2. request definition



```golang
type SysMenuForm struct {
	Id int64 `json:"id,optional"`
	ParentId int64 `json:"parentId"`
	Name string `json:"name"`
	Type int64 `json:"type,optional"`
	Path string `json:"path,optional"`
	Component string `json:"component,optional"`
	Perm string `json:"perm,optional"`
	Visible int64 `json:"visible,optional"`
	Sort int64 `json:"sort,default=false"`
	Icon string `json:"icon,optional"`
	Redirect string `json:"redirect,optional"`
}
```


3. response definition



```golang
type IdResp struct {
	ID int64 `json:"id"`
}
```

### 18. "删除菜单"

1. route definition

- Url: /sys/menu/del/:ids
- Method: GET
- Request: `PathIdsReq`
- Response: `-`

2. request definition



```golang
type PathIdsReq struct {
	Ids string `path:"ids"`
}
```


3. response definition


### 19. "角色分页列表"

1. route definition

- Url: /sys/role/list
- Method: GET
- Request: `SysRoleListReq`
- Response: `SysRoleListResp`

2. request definition



```golang
type SysRoleListReq struct {
	Keywords string `form:"keywords,omitempty,optional"`
	PageNum int64 `form:"pageNum,omitempty,optional,default=1"`
	PageSize int64 `form:"pageSize,omitempty,optional,default=10"`
}
```


3. response definition



```golang
type SysRoleListResp struct {
	List []*SysRoleInfo `json:"list"`
	Total int64 `json:"total"`
}
```

### 20. "新增角色"

1. route definition

- Url: /sys/role/add
- Method: POST
- Request: `SysRoleFormReq`
- Response: `IdResp`

2. request definition



```golang
type SysRoleFormReq struct {
	Id int64 `json:"id,optional"`
	Name string `json:"name"` // 角色名称
	Code string `json:"code"` // 角色编码
	Sort int64 `json:"sort"` // 显示顺序
	Status int64 `json:"status"` // 角色状态(1-正常；0-停用)
	DataScope int64 `json:"dataScope"` // 数据权限(0-所有数据；1-部门及子部门数据；2-本部门数据；3-本人数据)
}
```


3. response definition



```golang
type IdResp struct {
	ID int64 `json:"id"`
}
```

### 21. "获取角色"

1. route definition

- Url: /sys/role/:id
- Method: GET
- Request: `PathIdReq`
- Response: `SysRoleInfo`

2. request definition



```golang
type PathIdReq struct {
	Id int64 `path:"id"`
}
```


3. response definition



```golang
type SysRoleInfo struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	Sort int64 `json:"sort"`
	Status int64 `json:"status"`
	DataScope int64 `json:"dataScope"`
}
```

### 22. "角色下拉列表"

1. route definition

- Url: /sys/role/options
- Method: GET
- Request: `-`
- Response: `SysRoleOptionsResp`

2. request definition



3. response definition



```golang
type SysRoleOptionsResp struct {
	List []*SysRoleOption `json:"list"`
}
```

### 23. "修改角色"

1. route definition

- Url: /sys/role/update
- Method: POST
- Request: `SysRoleFormReq`
- Response: `IdResp`

2. request definition



```golang
type SysRoleFormReq struct {
	Id int64 `json:"id,optional"`
	Name string `json:"name"` // 角色名称
	Code string `json:"code"` // 角色编码
	Sort int64 `json:"sort"` // 显示顺序
	Status int64 `json:"status"` // 角色状态(1-正常；0-停用)
	DataScope int64 `json:"dataScope"` // 数据权限(0-所有数据；1-部门及子部门数据；2-本部门数据；3-本人数据)
}
```


3. response definition



```golang
type IdResp struct {
	ID int64 `json:"id"`
}
```

### 24. "删除角色"

1. route definition

- Url: /sys/role/del/:ids
- Method: GET
- Request: `PathIdsReq`
- Response: `-`

2. request definition



```golang
type PathIdsReq struct {
	Ids string `path:"ids"`
}
```


3. response definition


### 25. "获取角色的菜单ID集合"

1. route definition

- Url: /sys/role/menuIds/:id
- Method: GET
- Request: `PathIdReq`
- Response: `SysMenuIdsResp`

2. request definition



```golang
type PathIdReq struct {
	Id int64 `path:"id"`
}
```


3. response definition



```golang
type SysMenuIdsResp struct {
	Items []int64 `json:"items"`
}
```

### 26. "分配菜单权限给角色"

1. route definition

- Url: /sys/role/setMenuIds
- Method: POST
- Request: `SysSetMenuIdsReq`
- Response: `-`

2. request definition



```golang
type SysSetMenuIdsReq struct {
	RoleId int64 `json:"roleId,optional"`
	MenuIds []int64 `json:"menuIds,optional"`
}
```


3. response definition


### 27. "修改用户状态"

1. route definition

- Url: /sys/role/changeStatus
- Method: POST
- Request: `SysRoleChangeStatusReq`
- Response: `-`

2. request definition



```golang
type SysRoleChangeStatusReq struct {
	RoleId int64 `json:"roleId"`
	Status string `json:"status"`
}
```


3. response definition


### 28. "部门分页列表"

1. route definition

- Url: /sys/dept/list
- Method: GET
- Request: `SysDeptListReq`
- Response: `SysDeptListResp`

2. request definition



```golang
type SysDeptListReq struct {
	Keywords string `json:"keywords,omitempty,optional"`
	Status int64 `json:"status,omitempty,optional,default=-1"`
}
```


3. response definition



```golang
type SysDeptListResp struct {
	List []*SysDeptItem `json:"list"`
}
```

### 29. "新增部门"

1. route definition

- Url: /sys/dept/add
- Method: POST
- Request: `SysDeptFormReq`
- Response: `IdResp`

2. request definition



```golang
type SysDeptFormReq struct {
	Id int64 `json:"id,optional"`
	Name string `json:"name"`
	ParentId int64 `json:"parent_id"`
	TreePath string `json:"tree_path"`
	Sort int64 `json:"sort"`
	Status int64 `json:"status"`
	CreateBy int64 `json:"create_by"`
	UpdateBy int64 `json:"update_by"`
}
```


3. response definition



```golang
type IdResp struct {
	ID int64 `json:"id"`
}
```

### 30. "获取部门"

1. route definition

- Url: /sys/dept/:id
- Method: GET
- Request: `PathIdReq`
- Response: `SysDeptInfo`

2. request definition



```golang
type PathIdReq struct {
	Id int64 `path:"id"`
}
```


3. response definition



```golang
type SysDeptInfo struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	ParentId int64 `json:"parent_id"`
	TreePath string `json:"tree_path"`
	Sort int64 `json:"sort"`
	Status int64 `json:"status"`
	CreateBy int64 `json:"create_by"`
	UpdateBy int64 `json:"update_by"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
```

### 31. "部门下拉列表"

1. route definition

- Url: /sys/dept/options
- Method: GET
- Request: `-`
- Response: `SysDeptOptionsResp`

2. request definition



3. response definition



```golang
type SysDeptOptionsResp struct {
	List []*SysDeptOption `json:"list"`
}
```

### 32. "修改部门"

1. route definition

- Url: /sys/dept/update
- Method: POST
- Request: `SysDeptFormReq`
- Response: `IdResp`

2. request definition



```golang
type SysDeptFormReq struct {
	Id int64 `json:"id,optional"`
	Name string `json:"name"`
	ParentId int64 `json:"parent_id"`
	TreePath string `json:"tree_path"`
	Sort int64 `json:"sort"`
	Status int64 `json:"status"`
	CreateBy int64 `json:"create_by"`
	UpdateBy int64 `json:"update_by"`
}
```


3. response definition



```golang
type IdResp struct {
	ID int64 `json:"id"`
}
```

### 33. "删除部门"

1. route definition

- Url: /sys/dept/del/:ids
- Method: GET
- Request: `PathIdsReq`
- Response: `-`

2. request definition



```golang
type PathIdsReq struct {
	Ids string `path:"ids"`
}
```


3. response definition


### 34. "TypeList"

1. route definition

- Url: /sys/dict/typeList
- Method: GET
- Request: `SysDictTypeListReq`
- Response: `SysDictTypeListResp`

2. request definition



```golang
type SysDictTypeListReq struct {
	Keywords string `form:"keywords,omitempty,optional"`
	PageNum int64 `form:"pageNum,omitempty,optional,default=1"`
	PageSize int64 `form:"pageSize,omitempty,optional,default=10"`
}
```


3. response definition



```golang
type SysDictTypeListResp struct {
	List []*SysDictTypeItem `json:"list"`
	Total int64 `json:"total"`
}
```

### 35. "List"

1. route definition

- Url: /sys/dict/list
- Method: GET
- Request: `SysDictListReq`
- Response: `SysDictListResp`

2. request definition



```golang
type SysDictListReq struct {
	TypeCode string `form:"typeCode"`
	PageNum int64 `form:"pageNum,omitempty,optional,default=1"`
	PageSize int64 `form:"pageSize,omitempty,optional,default=10"`
}
```


3. response definition



```golang
type SysDictListResp struct {
	List []*SysDictItem `json:"list"`
	Total int64 `json:"total"`
}
```

### 36. "Options"

1. route definition

- Url: /sys/dict/options
- Method: GET
- Request: `-`
- Response: `SysDictOptionsResp`

2. request definition



3. response definition



```golang
type SysDictOptionsResp struct {
	List []*SysDictOption `json:"list"`
}
```

### 37. "TypeOptions"

1. route definition

- Url: /sys/dict/:typeCode/options
- Method: GET
- Request: `SysTypeReq`
- Response: `SysDictOptionsResp`

2. request definition



```golang
type SysTypeReq struct {
	TypeCode string `path:"typeCode"`
}
```


3. response definition



```golang
type SysDictOptionsResp struct {
	List []*SysDictOption `json:"list"`
}
```

### 38. "Add"

1. route definition

- Url: /sys/dict/add
- Method: POST
- Request: `SysDictForm`
- Response: `IdResp`

2. request definition



```golang
type SysDictForm struct {
	Id int64 `json:"id,optional"`
	TypeCode string `json:"type_code"`
	Name string `json:"name"`
	Value string `json:"value"`
	Sort int64 `json:"sort"`
	Status int64 `json:"status"`
	Defaulted int64 `json:"defaulted,default=1"`
	Remark string `json:"remark,optional"`
}
```


3. response definition



```golang
type IdResp struct {
	ID int64 `json:"id"`
}
```

### 39. "Update"

1. route definition

- Url: /sys/dict/update
- Method: POST
- Request: `SysDictForm`
- Response: `-`

2. request definition



```golang
type SysDictForm struct {
	Id int64 `json:"id,optional"`
	TypeCode string `json:"type_code"`
	Name string `json:"name"`
	Value string `json:"value"`
	Sort int64 `json:"sort"`
	Status int64 `json:"status"`
	Defaulted int64 `json:"defaulted,default=1"`
	Remark string `json:"remark,optional"`
}
```


3. response definition


### 40. "Get"

1. route definition

- Url: /sys/dict/:id
- Method: GET
- Request: `PathIdReq`
- Response: `SysDictItem`

2. request definition



```golang
type PathIdReq struct {
	Id int64 `path:"id"`
}
```


3. response definition



```golang
type SysDictItem struct {
	Id int64 `json:"id"`
	TypeCode string `json:"type_code"`
	Name string `json:"name"`
	Value string `json:"value"`
	Sort int64 `json:"sort,default=1"`
	Status int64 `json:"status"`
	Defaulted int64 `json:"defaulted,default=1"`
	Remark string `json:"remark,optional"`
	CreatedAt string `json:"created_at,optional"`
	UpdatedAt string `json:"updated_at,optional"`
}
```

### 41. "Del"

1. route definition

- Url: /sys/dict/del/:ids
- Method: GET
- Request: `PathIdsReq`
- Response: `-`

2. request definition



```golang
type PathIdsReq struct {
	Ids string `path:"ids"`
}
```


3. response definition


### 42. "TypeAdd"

1. route definition

- Url: /sys/dict/type/add
- Method: POST
- Request: `SysDictTypeForm`
- Response: `IdResp`

2. request definition



```golang
type SysDictTypeForm struct {
	Id int64 `json:"id,optional"`
	Name string `json:"name"`
	Code string `json:"code"`
	Status int64 `json:"status"`
	Remark string `json:"remark,optional"`
}
```


3. response definition



```golang
type IdResp struct {
	ID int64 `json:"id"`
}
```

### 43. "TypeUpdate"

1. route definition

- Url: /sys/dict/type/update
- Method: POST
- Request: `SysDictTypeForm`
- Response: `-`

2. request definition



```golang
type SysDictTypeForm struct {
	Id int64 `json:"id,optional"`
	Name string `json:"name"`
	Code string `json:"code"`
	Status int64 `json:"status"`
	Remark string `json:"remark,optional"`
}
```


3. response definition


### 44. "TypeGet"

1. route definition

- Url: /sys/dict/type/:id
- Method: GET
- Request: `PathIdReq`
- Response: `SysDictTypeItem`

2. request definition



```golang
type PathIdReq struct {
	Id int64 `path:"id"`
}
```


3. response definition



```golang
type SysDictTypeItem struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	Status int64 `json:"status"`
	Remark string `json:"remark,optional"`
	CreatedAt string `json:"created_at,optional"`
	UpdatedAt string `json:"updated_at,optional"`
}
```

### 45. "TypeDel"

1. route definition

- Url: /sys/dict/type/del/:ids
- Method: GET
- Request: `PathIdsReq`
- Response: `-`

2. request definition



```golang
type PathIdsReq struct {
	Ids string `path:"ids"`
}
```


3. response definition


### 46. "添加用户"

1. route definition

- Url: /user/add
- Method: POST
- Request: `UserForm`
- Response: `IdResp`

2. request definition



```golang
type UserForm struct {
	Id int64 `json:"id,optional"`
	Mobile string `json:"mobile,optional"` //电话号码
	Username string `json:"username,optional"` //姓名
	Nickname string `json:"nickname,optional"` //昵称
	Avatar string `json:"avatar,optional"` //头像
	IdCard string `json:"idCard,optional"` //身份证号码
	Gender int64 `json:"gender,optional"` //性别 0 未知 1男 2女
	Signature string `json:"signature,optional"`
	Birthday string `json:"birthday,optional"`
	Tags string `json:"tags,optional"`
	Source string `json:"source,optional"`
	SourceUid int64 `json:"sourceUid,optional"`
}
```


3. response definition



```golang
type IdResp struct {
	ID int64 `json:"id"`
}
```

### 47. "修改用户信息"

1. route definition

- Url: /user/update
- Method: POST
- Request: `UserForm`
- Response: `-`

2. request definition



```golang
type UserForm struct {
	Id int64 `json:"id,optional"`
	Mobile string `json:"mobile,optional"` //电话号码
	Username string `json:"username,optional"` //姓名
	Nickname string `json:"nickname,optional"` //昵称
	Avatar string `json:"avatar,optional"` //头像
	IdCard string `json:"idCard,optional"` //身份证号码
	Gender int64 `json:"gender,optional"` //性别 0 未知 1男 2女
	Signature string `json:"signature,optional"`
	Birthday string `json:"birthday,optional"`
	Tags string `json:"tags,optional"`
	Source string `json:"source,optional"`
	SourceUid int64 `json:"sourceUid,optional"`
}
```


3. response definition


### 48. "获取用户信息"

1. route definition

- Url: /user/:id
- Method: GET
- Request: `PathIdReq`
- Response: `UserInfo`

2. request definition



```golang
type PathIdReq struct {
	Id int64 `path:"id"`
}
```


3. response definition



```golang
type UserInfo struct {
	Id int64 `json:"id"`
	Mobile string `json:"mobile"` //电话号码
	Username string `json:"username,optional"`
	Nickname string `json:"nickname"` //昵称
	Avatar string `json:"avatar"` //头像
	IdCard string `json:"idCard"` //身份证号码
	Gender int8 `json:"gender,optional"` //性别 0 未知 1男 2女
	Signature string `json:"signature"` //签名
	Birthday string `json:"birthday"` //生日
	Tags string `json:"tags"` //tags
	Source string `json:"source"` //来源，APP H5
	SourceUid int64 `json:"sourceUid"` //邀请uid
	CreatedAt string `json:"created_at"`
	Addresses []*UserAddress `json:"addresses"`
}
```

### 49. "获取用户列表"

1. route definition

- Url: /user/list
- Method: GET
- Request: `ListReq`
- Response: `UserListRsqp`

2. request definition



```golang
type ListReq struct {
	Keywords string `form:"keywords,omitempty,optional"`
	Status int64 `form:"status,default=-1"`
	PageNum int64 `form:"pageNum,default=1"`
	PageSize int64 `form:"pageSize,default=10"`
}
```


3. response definition



```golang
type UserListRsqp struct {
	List []*UserInfo `json:"list"`
	Total int64 `json:"total"`
}
```

### 50. "删除用户"

1. route definition

- Url: /user/del/:ids
- Method: GET
- Request: `PathIdsReq`
- Response: `-`

2. request definition



```golang
type PathIdsReq struct {
	Ids string `path:"ids"`
}
```


3. response definition


### 51. "获取第三方用户列表"

1. route definition

- Url: /user/thirdList
- Method: GET
- Request: `ListReq`
- Response: `UserThirdListRsqp`

2. request definition



```golang
type ListReq struct {
	Keywords string `form:"keywords,omitempty,optional"`
	Status int64 `form:"status,default=-1"`
	PageNum int64 `form:"pageNum,default=1"`
	PageSize int64 `form:"pageSize,default=10"`
}
```


3. response definition



```golang
type UserThirdListRsqp struct {
	List []*UserThird `json:"list"`
	Total int64 `json:"total"`
}
```

### 52. "新增类型"

1. route definition

- Url: /shop/add
- Method: POST
- Request: `ShopForm`
- Response: `IdResp`

2. request definition



```golang
type ShopForm struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Tel string `json:"tel"`
	Notice string `json:"notice"`
	DistributionFee float64 `json:"distributionFee"`
}
```


3. response definition



```golang
type IdResp struct {
	ID int64 `json:"id"`
}
```

### 53. "修改类型"

1. route definition

- Url: /shop/update
- Method: POST
- Request: `ShopForm`
- Response: `-`

2. request definition



```golang
type ShopForm struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Tel string `json:"tel"`
	Notice string `json:"notice"`
	DistributionFee float64 `json:"distributionFee"`
}
```


3. response definition


### 54. "店铺下拉列表"

1. route definition

- Url: /shop/options
- Method: GET
- Request: `-`
- Response: `OptionsResp`

2. request definition



3. response definition



```golang
type OptionsResp struct {
	List []*OptionItem `json:"list"`
}
```

### 55. "获取店铺"

1. route definition

- Url: /shop/:id
- Method: GET
- Request: `PathIdReq`
- Response: `Shop`

2. request definition



```golang
type PathIdReq struct {
	Id int64 `path:"id"`
}
```


3. response definition



```golang
type Shop struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Tel string `json:"tel"`
	Notice string `json:"notice"`
	DistributionFee float64 `json:"distributionFee"`
	Status int64 `json:"status"`
	CreatedAt string `json:"createdAt"`
}
```

### 56. "店铺列表"

1. route definition

- Url: /shop/list
- Method: GET
- Request: `ListReq`
- Response: `ShopListResp`

2. request definition



```golang
type ListReq struct {
	Keywords string `form:"keywords,omitempty,optional"`
	Status int64 `form:"status,default=-1"`
	PageNum int64 `form:"pageNum,default=1"`
	PageSize int64 `form:"pageSize,default=10"`
}
```


3. response definition



```golang
type ShopListResp struct {
	List []*Shop `json:"list"`
	Total int64 `json:"total"`
}
```

### 57. "删除"

1. route definition

- Url: /shop/del/:ids
- Method: GET
- Request: `PathIdsReq`
- Response: `-`

2. request definition



```golang
type PathIdsReq struct {
	Ids string `path:"ids"`
}
```


3. response definition


### 58. "新增类型"

1. route definition

- Url: /shop/category/add
- Method: POST
- Request: `PmsCategoryForm`
- Response: `IdResp`

2. request definition



```golang
type PmsCategoryForm struct {
	Id int64 `json:"id"`
	ShopId int64 `json:"shopId"`
	ParentId int64 `json:"parentId"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Sort int64 `json:"sort"`
	Visible int64 `json:"visible"`
}
```


3. response definition



```golang
type IdResp struct {
	ID int64 `json:"id"`
}
```

### 59. "修改类型"

1. route definition

- Url: /shop/category/update
- Method: POST
- Request: `PmsCategoryForm`
- Response: `-`

2. request definition



```golang
type PmsCategoryForm struct {
	Id int64 `json:"id"`
	ShopId int64 `json:"shopId"`
	ParentId int64 `json:"parentId"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Sort int64 `json:"sort"`
	Visible int64 `json:"visible"`
}
```


3. response definition


### 60. "类型下拉列表"

1. route definition

- Url: /shop/category/options/:id
- Method: GET
- Request: `PathIdReq`
- Response: `OptionsResp`

2. request definition



```golang
type PathIdReq struct {
	Id int64 `path:"id"`
}
```


3. response definition



```golang
type OptionsResp struct {
	List []*OptionItem `json:"list"`
}
```

### 61. "获取店铺"

1. route definition

- Url: /shop/category/:id
- Method: GET
- Request: `PathIdReq`
- Response: `PmsCategory`

2. request definition



```golang
type PathIdReq struct {
	Id int64 `path:"id"`
}
```


3. response definition



```golang
type PmsCategory struct {
	Id int64 `json:"id"`
	ShopId int64 `json:"shopId"`
	ParentId int64 `json:"parentId"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Sort int64 `json:"sort"`
	Visible int64 `json:"visible"`
	CreatedAt string `json:"createdAt"`
	Children []*PmsCategory `json:"children,omitempty"`
}
```

### 62. "类型列表"

1. route definition

- Url: /shop/category/list
- Method: GET
- Request: `ListReq`
- Response: `PmsCategoryListResp`

2. request definition



```golang
type ListReq struct {
	Keywords string `form:"keywords,omitempty,optional"`
	Status int64 `form:"status,default=-1"`
	PageNum int64 `form:"pageNum,default=1"`
	PageSize int64 `form:"pageSize,default=10"`
}
```


3. response definition



```golang
type PmsCategoryListResp struct {
	List []*PmsCategory `json:"list"`
}
```

### 63. "删除类型"

1. route definition

- Url: /shop/category/del/:ids
- Method: GET
- Request: `PathIdsReq`
- Response: `-`

2. request definition



```golang
type PathIdsReq struct {
	Ids string `path:"ids"`
}
```


3. response definition


### 64. "新增类型-属性"

1. route definition

- Url: /shop/category/attribute/add
- Method: POST
- Request: `PmsCategoryAttributeForm`
- Response: `-`

2. request definition



```golang
type PmsCategoryAttributeForm struct {
	CategoryId int64 `json:"categoryId"` // 分类ID
	Attributes []string `json:"attributes"` // 属性名称
	Type int64 `json:"type"` // 类型(1:规格;2:属性;)
}
```


3. response definition


### 65. "新增类型-属性"

1. route definition

- Url: /shop/category/attribute/list
- Method: GET
- Request: `PmsCategoryAttributeListReq`
- Response: `PmsCategoryAttributeListResp`

2. request definition



```golang
type PmsCategoryAttributeListReq struct {
	CategoryId int64 `form:"categoryId"` // 分类ID
	Type int64 `form:"type"` // 类型(1:规格;2:属性;)
}
```


3. response definition



```golang
type PmsCategoryAttributeListResp struct {
	List []*PmsCategoryAttribute `json:"list"`
}
```

### 66. "新增Brand"

1. route definition

- Url: /shop/brand/add
- Method: POST
- Request: `BrandForm`
- Response: `IdResp`

2. request definition



```golang
type BrandForm struct {
	Id int64 `json:"id"` // 主键
	ShopId int64 `json:"shopId"` // shopID
	Name string `json:"name"` // 品牌名称
	Logo string `json:"logo"` // LOGO图片
	Sort int64 `json:"sort"` // 排序
}
```


3. response definition



```golang
type IdResp struct {
	ID int64 `json:"id"`
}
```

### 67. "修改Brand"

1. route definition

- Url: /shop/brand/update
- Method: POST
- Request: `BrandForm`
- Response: `-`

2. request definition



```golang
type BrandForm struct {
	Id int64 `json:"id"` // 主键
	ShopId int64 `json:"shopId"` // shopID
	Name string `json:"name"` // 品牌名称
	Logo string `json:"logo"` // LOGO图片
	Sort int64 `json:"sort"` // 排序
}
```


3. response definition


### 68. "Brand下拉列表"

1. route definition

- Url: /shop/brand/options/:id
- Method: GET
- Request: `PathIdReq`
- Response: `OptionsResp`

2. request definition



```golang
type PathIdReq struct {
	Id int64 `path:"id"`
}
```


3. response definition



```golang
type OptionsResp struct {
	List []*OptionItem `json:"list"`
}
```

### 69. "获取Brand"

1. route definition

- Url: /shop/brand/:id
- Method: GET
- Request: `PathIdReq`
- Response: `PmsBrand`

2. request definition



```golang
type PathIdReq struct {
	Id int64 `path:"id"`
}
```


3. response definition



```golang
type PmsBrand struct {
	Id int64 `json:"id"` // 主键
	ShopId int64 `json:"shopId"` // shopID
	Name string `json:"name"` // 品牌名称
	Logo string `json:"logo"` // LOGO图片
	Sort int64 `json:"sort"` // 排序
	CreatedAt string `json:"createdAt"`
}
```

### 70. "Brand列表"

1. route definition

- Url: /shop/brand/list
- Method: GET
- Request: `BrandListReq`
- Response: `BrandListResp`

2. request definition



```golang
type BrandListReq struct {
	Keywords string `form:"keywords,omitempty,optional"`
	ShopId int64 `json:"shopId,default=1"` // shopID
	PageNum int64 `form:"pageNum,default=1"`
	PageSize int64 `form:"pageSize,default=10"`
}
```


3. response definition



```golang
type BrandListResp struct {
	List []PmsBrand `json:"list"`
	Total int64 `json:"total"`
}
```

### 71. "删除"

1. route definition

- Url: /shop/brand/del/:ids
- Method: GET
- Request: `PathIdsReq`
- Response: `-`

2. request definition



```golang
type PathIdsReq struct {
	Ids string `path:"ids"`
}
```


3. response definition


### 72. "获取商品"

1. route definition

- Url: /goods/:id
- Method: GET
- Request: `PathIdReq`
- Response: `PmsGoods`

2. request definition



```golang
type PathIdReq struct {
	Id int64 `path:"id"`
}
```


3. response definition



```golang
type PmsGoods struct {
	Id int64 `json:"id"` // 主键
	ShopId int64 `json:"shopId"` // shopID
	CategoryId int64 `json:"categoryId"` // 商品类型ID
	BrandId int64 `json:"brandId"` // 商品品牌ID
	Name string `json:"name"` // 商品名称
	OriginPrice int64 `json:"originPrice"` // 原价【起】
	Price int64 `json:"price"` // 现价【起】
	Sales int64 `json:"sales"` // 销量
	PicUrl string `json:"picUrl"` // 商品主图
	SubPicUrls []string `json:"subPicUrls"` // 商品图册
	Unit string `json:"unit"` // 单位
	Description string `json:"description"` // 商品简介
	Detail string `json:"detail"` // 商品详情
	Status int64 `json:"status"` // 商品状态(0:下架 1:上架)
	BrandName string `json:"brandName,omitempty,optional"`
	CategoryName string `json:"categoryName,omitempty,optional"`
	SkuList []PmsSku `json:"skuList,omitempty,optional"`
	AttrList []PmsGoodsAttribute `json:"attrList,omitempty,optional"`
	SpecList []PmsGoodsAttribute `json:"specList,omitempty,optional"`
	CreatedAt string `json:"createdAt"`
}
```

### 73. "商品列表"

1. route definition

- Url: /goods/list
- Method: GET
- Request: `PmsGoodsListReq`
- Response: `PmsGoodsListRsqp`

2. request definition



```golang
type PmsGoodsListReq struct {
	ShopId int64 `form:"shopId,default=-1"`
	Keywords string `form:"keywords,omitempty,optional"`
	Status int64 `form:"status,default=-1"`
	PageNum int64 `form:"pageNum,default=1"`
	PageSize int64 `form:"pageSize,default=10"`
	CategoryId int64 `form:"categoryId,omitempty,optional"`
}
```


3. response definition



```golang
type PmsGoodsListRsqp struct {
	List []PmsGoods `json:"list"`
	Total int64 `json:"total"`
}
```

### 74. "新增商品"

1. route definition

- Url: /goods/edit
- Method: POST
- Request: `GoodsEditRep`
- Response: `IdResp`

2. request definition



```golang
type GoodsEditRep struct {
	Id int64 `json:"id,optional"`
	ShopId int64 `json:"shopId"`
	CategoryId int64 `json:"categoryId"`
	BrandId int64 `json:"brandId"`
	Name string `json:"name"`
	OriginPrice int64 `json:"originPrice"`
	Price int64 `json:"price"`
	PicUrl string `json:"picUrl,optional"`
	SubPicUrls []string `json:"subPicUrls,optional"`
	Unit string `json:"unit,default=个"`
	Description string `json:"description"`
	Detail string `json:"detail"`
	Status int64 `json:"status,default=1"`
	SkuList []PmsGoodsFormSku `json:"skuList,optional"`
	AttrList []PmsGoodsFormAttribute `json:"attrList,optional"`
	SpecList []PmsGoodsFormAttribute `json:"specList,optional"`
}
```


3. response definition



```golang
type IdResp struct {
	ID int64 `json:"id"`
}
```

### 75. "删除"

1. route definition

- Url: /goods/del/:ids
- Method: GET
- Request: `PathIdsReq`
- Response: `-`

2. request definition



```golang
type PathIdsReq struct {
	Ids string `path:"ids"`
}
```


3. response definition


### 76. "订单列表"

1. route definition

- Url: /order/list
- Method: GET
- Request: `ListReq`
- Response: `OrderListRsqp`

2. request definition



```golang
type ListReq struct {
	Keywords string `form:"keywords,omitempty,optional"`
	Status int64 `form:"status,default=-1"`
	PageNum int64 `form:"pageNum,default=1"`
	PageSize int64 `form:"pageSize,default=10"`
}
```


3. response definition



```golang
type OrderListRsqp struct {
	List []*OrderInfo `json:"list"`
	Total int64 `json:"total"`
}
```

### 77. "获取订单"

1. route definition

- Url: /order/:id
- Method: GET
- Request: `PathIdReq`
- Response: `OrderInfo`

2. request definition



```golang
type PathIdReq struct {
	Id int64 `path:"id"`
}
```


3. response definition



```golang
type OrderInfo struct {
	Id int64 `json:"id"`
	ShopId int64 `json:"shop_id"` // 店铺ID
	UserId int64 `json:"user_id"` // 用户ID
	OrderNo string `json:"order_no"` // 订单号
	GoodsNumber int64 `json:"goods_number"` // 商品总数
	PayAmount float64 `json:"pay_amount"` // 订单金额(商品价格&#43;配送费-优惠金额)
	GoodsAmount float64 `json:"goods_amount"` // 商品费用
	FreightAmount float64 `json:"freight_amount"` // 运费
	CouponAmount float64 `json:"coupon_amount"` // 优惠金费
	CouponId int64 `json:"coupon_id"` // 用户优惠券id
	ReceiverType int64 `json:"receiver_type"` // 地址类型0 配送 1自提点
	ReceiverName string `json:"receiver_name"` // 收货人姓名
	ReceiverMobile string `json:"receiver_mobile"` // 收货人电话
	ReceiverAddress string `json:"receiver_address"` // 收货人地址
	DistributionId int64 `json:"distribution_id"` // 自提点ID
	DistributionInfo string `json:"distribution_info"` // 自提点ID详情
	Remarks string `json:"remarks"` // 备注
	PayTime int64 `json:"pay_time"` // 支付时间
	TradeNo string `json:"trade_no"` // 支付单号
	Status int64 `json:"status"` // 订单状态  -2退款 -1=已关闭/0待支付/1已支付/2已发货/5待取货/10完成
	AppraiseStatus int64 `json:"appraise_status"` // 评价状态：0.待评价；1.已平价;
	CreatedAt string `json:"created_at"`
}
```

### 78. "文件上传"

1. route definition

- Url: /other/upload
- Method: POST
- Request: `-`
- Response: `FileUpload`

2. request definition



3. response definition



```golang
type FileUpload struct {
	Hash string `json:"hash,optional"`
	Name string `json:"name,optional"`
	Ext string `json:"ext,optional"`
	Size int64 `json:"size,optional"`
	Path string `json:"path,optional"`
}
```

