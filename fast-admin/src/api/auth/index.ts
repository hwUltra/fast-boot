import request from "@/utils/request";

const AUTH_BASE_URL = "";

const AuthAPI = {
  /** 登录 接口*/
  login(data: LoginData) {
    return request<any, LoginResult>({
      url: `${AUTH_BASE_URL}/auth/login`,
      method: "post",
      data: data,
      // headers: {
      //   "Content-Type": "multipart/form-data",
      // },
    });
  },

  /** 注销 接口*/
  logout() {
    return request({
      url: `${AUTH_BASE_URL}/auth/logout`,
      method: "get",
    });
  },

  /** 获取验证码 接口*/
  getCaptcha() {
    return request<any, CaptchaResult>({
      url: `${AUTH_BASE_URL}/other/captcha`,
      method: "get",
    });
  },

  /**
   * 获取当前登录用户信息
   *
   * @returns 登录用户昵称、头像信息，包括角色和权限
   */
  getInfo() {
    return request<any, UserInfo>({
      url: `${AUTH_BASE_URL}/auth/me`,
      method: "get",
    });
  },
};

export default AuthAPI;

/** 登录请求参数 */
export interface LoginData {
  /** 用户名 */
  username: string;
  /** 密码 */
  password: string;
  /** 验证码缓存key */
  captchaKey: string;
  /** 验证码 */
  captchaCode: string;
}

/** 登录响应 */
export interface LoginResult {
  /** 访问token */
  accessToken?: string;
  /** 过期时间(单位：毫秒) */
  expires?: number;
  /** 刷新token */
  refreshToken?: string;
  /** token 类型 */
  tokenType?: string;
}

/** 验证码响应 */
export interface CaptchaResult {
  /** 验证码缓存key */
  captchaKey: string;
  /** 验证码图片Base64字符串 */
  captchaBase64: string;
}

/** 登录用户信息 */
export interface UserInfo {
  /** 用户ID */
  userId?: number;

  /** 用户名 */
  username?: string;

  /** 昵称 */
  nickname?: string;

  /** 头像URL */
  avatar?: string;

  /** 角色 */
  roles: string[];

  /** 权限 */
  perms: string[];
}
