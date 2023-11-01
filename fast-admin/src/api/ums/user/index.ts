import request from "@/utils/request";
import { AxiosPromise } from "axios";
import { UserForm, UserPageVO, UserQuery, UserThirdPageVO } from "./types";

/**
 * 获取用户分页列表
 *
 * @param queryParams
 */
export function getList(
  queryParams: UserQuery
): AxiosPromise<PageResult<UserPageVO[]>> {
  return request({
    url: "/user/list",
    method: "get",
    params: queryParams,
  });
}

/**
 * 获取用户表单详情
 *
 * @param userId
 */
export function getUserInfo(userId: number): AxiosPromise<UserForm> {
  return request({
    url: "/user/" + userId,
    method: "get",
  });
}

/**
 * 添加用户
 *
 * @param data
 */
export function addUser(data: UserForm) {
  return request({
    url: "/user/add",
    method: "post",
    data: data,
  });
}

/**
 * 修改用户
 *
 * @param id
 * @param data
 */
export function updateUser(data: UserForm) {
  // data.gender = Number(data.gender);
  return request({
    url: "/user/update",
    method: "post",
    data: data,
  });
}

/**
 * 删除用户
 *
 * @param ids
 */
export function delUser(ids: string) {
  return request({
    url: "/user/del/" + ids,
    method: "get",
  });
}

/**
 * 获取用户分页列表
 *
 * @param queryParams
 */
export function getThirdList(
  queryParams: UserQuery
): AxiosPromise<PageResult<UserThirdPageVO[]>> {
  return request({
    url: "/user/thirdList",
    method: "get",
    params: queryParams,
  });
}
