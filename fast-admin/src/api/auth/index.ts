import request from "@/utils/request";
import { AxiosPromise } from "axios";
import { LoginData, LoginResult } from "./types";

/**
 * 登录API
 *
 * @param data {LoginData}
 * @returns
 */
export function loginApi(data: LoginData): AxiosPromise<LoginResult> {
  // const formData = new FormData();
  // formData.append("username", data.username);
  // formData.append("password", data.password);
  // formData.append("captchaKey", data.captchaKey || "");
  // formData.append("captchaCode", data.captchaCode || "");
  return request({
    url: "/auth/login",
    method: "post",
    data: data,
  });
}

/**
 * 注销API
 */
export function logoutApi() {
  return request({
    url: "/auth/logout",
    method: "get",
  });
}

/**
 * 获取验证码
 */
export function getCaptchaApi(): AxiosPromise<CaptchaResult> {
  return request({
    url: "/other/captcha",
    method: "get",
  });
}
