import request from "@/utils/request";
import { AxiosPromise } from "axios";
import { ShopForm, ShopQuery, ShopVO } from "./types";

export function shopInfo(id: number): AxiosPromise<ShopVO> {
  return request({
    url: "/shop/" + id,
    method: "get",
  });
}

export function shopList(
  queryParams?: ShopQuery
): AxiosPromise<PageResult<ShopVO[]>> {
  return request({
    url: "/shop/list",
    method: "get",
    params: queryParams,
  });
}

export function getShopOptions(): AxiosPromise<ShopOptions> {
  return request({
    url: "/shop/options",
    method: "get",
  });
}

export function shopAdd(data: ShopForm) {
  return request({
    url: "/shop/add",
    method: "post",
    data: data,
  });
}

export function shopUpdate(data: ShopForm) {
  return request({
    url: "/shop/update",
    method: "post",
    data: data,
  });
}

export function shopDel(ids: string) {
  return request({
    url: "/shop/del/" + ids,
    method: "get",
  });
}
