import request from "@/utils/request";
import { AxiosPromise } from "axios";
import { DeptForm, DeptQuery, DeptVO } from "./types";

/**
 * 部门树形表格
 *
 * @param queryParams
 */
export function listDepts(queryParams?: DeptQuery): AxiosPromise<DeptVO[]> {
  return request({
    url: "/sys/dept/list",
    method: "get",
    params: queryParams,
  });
}

/**
 * 部门下拉列表
 */
export function getDeptOptions(): AxiosPromise<[]> {
  return request({
    url: "/sys/dept/options",
    method: "get",
  });
}

/**
 * 获取部门详情
 *
 * @param id
 */
export function getDeptForm(id: number): AxiosPromise<DeptForm> {
  return request({
    url: "/sys/dept/" + id,
    method: "get",
  });
}

/**
 * 新增部门
 *
 * @param data
 */
export function addDept(data: DeptForm) {
  return request({
    url: "/sys/dept/add",
    method: "post",
    data: data,
  });
}

/**
 *  修改部门
 *
 * @param id
 * @param data
 */
export function updateDept(id: number, data: DeptForm) {
  const deptData = { ...data };
  deptData.id = id;
  return request({
    url: "/sys/dept/update",
    method: "post",
    data: deptData,
  });
}

/**
 * 删除部门
 *
 * @param ids
 */
export function deleteDept(ids: string) {
  return request({
    url: "/sys/dept/del/" + ids,
    method: "get",
  });
}
