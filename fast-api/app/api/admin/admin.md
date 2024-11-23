### 1. 登录

1. route definition

- Url: /auth/login
- Method: POST
- Request: `LoginReq`
- Response: `TokenResp`

2. request definition



```golang
type LoginReq struct {
	UserName string `json:"username" validate:"min=2,max=50"    label:"用户名"`
	Password string `json:"password" validate:"min=2,max=50"    label:"密码"`
	CaptchaCode string `json:"captchaCode" validate:"min=1,max=50"    label:"CaptchaCode"`
	CaptchaKey string `json:"captchaKey" validate:"min=2,max=50"    label:"CaptchaKey"`
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

### 2. 退出登录

1. route definition

- Url: /auth/logout
- Method: GET
- Request: `-`
- Response: `LogoutResp`

2. request definition



3. response definition



```golang
type LogoutResp struct {
}
```

### 3. 获取个人信息

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

### 4. 刷新ToKen

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

### 5. 验证码

1. route definition

- Url: /other/captcha
- Method: GET
- Request: `-`
- Response: `CaptchaResp`

2. request definition



3. response definition



```golang
type CaptchaResp struct {
	CaptchaBase64 string `json:"captchaBase64"`
	CaptchaKey string `json:"captchaKey"`
}
```

### 6. 文件上传

1. route definition

- Url: /other/upload
- Method: POST
- Request: `-`
- Response: `UploadInfo`

2. request definition



3. response definition



```golang
type UploadInfo struct {
	Name string `json:"name,optional"`
	Url string `json:"url,optional"`
}
```

