import request from "@/utils/request";
import { getToken } from "@/utils/cache";
import { ResultCodeEnum } from "@/enums/ResultCodeEnum";

// H5 使用 VITE_APP_BASE_API 作为代理路径，其他平台使用 VITE_APP_API_URL 作为请求路径
let baseApi = import.meta.env.VITE_APP_API_URL;
// #ifdef H5
baseApi = import.meta.env.VITE_APP_BASE_API;

const OtherAPI = {
  /**
   * 发送验证码
   *
   * @param username 用户名
   * @param password 密码
   * @returns 返回 token
   */
  sendSms(type: string, mobile: string): Promise<void> {
    return request({
      url: `/sendSms/${type}/${mobile}`,
      method: "GET",
    });
  },

  /**
   * 文件上传地址
   */
  uploadUrl: baseApi + "/upload",

  /**
   * 上传文件
   *
   * @param filePath
   */
  upload(filePath: string): Promise<FileInfo> {
    return new Promise((resolve, reject) => {
      uni.uploadFile({
        url: this.uploadUrl,
        filePath: filePath,
        name: "file",
        header: {
          Authorization: getToken() ? `Bearer ${getToken()}` : "",
        },
        formData: {},
        success: (response) => {
          const resData = JSON.parse(response.data) as ResponseData<FileInfo>;
          if (resData.code === ResultCodeEnum.SUCCESS) {
            resolve(resData.data);
          } else {
            // 其他业务处理失败
            uni.showToast({
              title: resData.msg || "文件上传失败",
              icon: "none",
            });
            reject({
              message: resData.msg || "业务处理失败",
              code: resData.code,
            });
          }
        },
        fail: (error) => {
          console.log("fail error", error);
          uni.showToast({
            title: "文件上传请求失败",
            icon: "none",
            duration: 2000,
          });
          reject({
            message: "文件上传请求失败",
            error,
          });
        },
      });
    });
  },
};

export default OtherAPI;

/**
 * 文件API类型声明
 */
export interface FileInfo {
  /** 文件名 */
  name: string;
  /** 文件路径 */
  url: string;
}
