import request from "@/utils/request";

const AuthAPI = {
  /**
   * 登录接口
   *
   * @param username 用户名
   * @param password 密码
   * @returns 返回 token
   */
  login(data: LoginFormData): Promise<LoginResult> {
    console.log("data", data);
    return request<LoginResult>({
      url: "/auth/login",
      method: "POST",
      data: data,
      // header: {
      //   "Content-Type": "application/x-www-form-urlencoded",
      // },
    });
  },

  /**
   * 登出接口
   */
  logout(): Promise<void> {
    return request({
      url: "/auth/logout",
      method: "GET",
    });
  },
};

export default AuthAPI;

/** 登录响应 */
export interface LoginResult {
  /** 访问token */
  accessToken: string;
}

export interface LoginFormData {
  mobile: string;
  vcode: string;
}
