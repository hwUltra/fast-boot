import request from "@/utils/request";
import { AxiosPromise } from "axios";
import { CategoryQuery, CategoryVO, CategoryForm } from "./types";

export function categoryList(
  queryParams?: CategoryQuery
): AxiosPromise<ListResult<CategoryVO[]>> {
  return request({
    url: "/shop/category/list",
    method: "get",
    params: queryParams,
  });
}

export function categoryInfo(id: number): AxiosPromise<CategoryVO> {
  return request({
    url: "/shop/category/" + id,
    method: "get",
  });
}

export function categoryAdd(data: CategoryForm) {
  return request({
    url: "/shop/category/add",
    method: "post",
    data: data,
  });
}

export function categoryUpdate(data: CategoryForm) {
  return request({
    url: "/shop/category/update",
    method: "post",
    data: data,
  });
}

export function categoryDel(ids: string) {
  return request({
    url: "/shop/category/del/" + ids,
    method: "get",
  });
}

/**
 * 部门下拉列表
 */
export function getOptions(id: number): AxiosPromise<CategoryOptions> {
  return request({
    url: "/shop/category/options/" + id,
    method: "get",
  });
}
