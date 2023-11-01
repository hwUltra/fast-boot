import request from "@/utils/request";
import { AxiosPromise } from "axios";
import { MenuQuery, MenuVO, MenuForm } from "./types";

/**
 * 获取路由列表
 */
export function listRoutes() {
  return request({
    url: "/sys/menu/routes",
    method: "get",
  });
}

/**
 * 获取菜单树形列表
 *
 * @param queryParams
 */
export function listMenus(queryParams: MenuQuery): AxiosPromise<MenuVO[]> {
  return request({
    url: "/sys/menu/list",
    method: "get",
    params: queryParams,
  });
}

/**
 * 获取菜单下拉树形列表
 */
export function getMenuOptions(): AxiosPromise<OptionType[]> {
  return request({
    url: "/sys/menu/options",
    method: "get",
  });
}

/**
 * 获取菜单表单数据
 *
 * @param id
 */
export function getMenuForm(id: number): AxiosPromise<MenuForm> {
  return request({
    url: "/sys/menu/" + id,
    method: "get",
  });
}

/**
 * 添加菜单
 *
 * @param data
 */
export function addMenu(data: MenuForm) {
  return request({
    url: "/sys/menu/add",
    method: "post",
    data: data,
  });
}

/**
 * 修改菜单
 *
 * @param id
 * @param data
 */
export function updateMenu(id: number, data: MenuForm) {
  const menuData = { ...data };
  menuData.id = id;
  return request({
    url: "/sys/menu/update",
    method: "post",
    data: menuData,
  });
}

/**
 * 删除菜单
 *
 * @param id 菜单ID
 */
export function deleteMenu(id: number) {
  return request({
    url: "/sys/menu/del/" + id,
    method: "get",
  });
}
