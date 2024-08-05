import { BrandForm, BrandPageResult, BrandQuery } from "./types";
import request from "@/utils/request";
import { AxiosPromise } from "axios";

export function getBrandList(
  queryParams?: BrandQuery
): AxiosPromise<BrandPageResult> {
  return request({
    url: "/shop/brand/list",
    params: queryParams,
    method: "get",
  });
}

/**
 * 获取品牌列表
 *
 * @param queryParams
 */
export function getBrandOptions(id: number): AxiosPromise<OptionType[]> {
  return request({
    url: "/shop/brand/options/" + id,
    method: "get",
  });
}

/**
 * 添加品牌
 *
 * @param data
 */
export function brandAdd(data: BrandForm) {
  return request({
    url: "/shop/brand/add",
    method: "post",
    data: data,
  });
}

/**
 * 修改品牌
 *
 * @param id
 * @param data
 */
export function brandUpdate(data: BrandForm) {
  return request({
    url: "/shop/brand/update",
    method: "post",
    data: data,
  });
}

/**
 * 删除品牌
 *
 * @param ids
 */
export function brandDel(ids: string) {
  return request({
    url: "/shop/brand/del/" + ids,
    method: "get",
  });
}
