import request from "@/utils/request";
const SHOP_BASE_URL = "/shop";

class ShopAPI {
  static getShopInfo(id: number) {
    return request<any, ShopVO>({
      url: `${SHOP_BASE_URL}/${id}`,
      method: "get",
    });
  }

  static getShopList(queryParams?: ShopQuery) {
    return request<any, PageResult<ShopVO[]>>({
      url: `${SHOP_BASE_URL}`,
      method: "get",
      params: queryParams,
    });
  }

  static getShopOptions() {
    return request<any, OptionType[]>({
      url: `${SHOP_BASE_URL}/options`,
      method: "get",
    });
  }

  static shopAdd(data: ShopForm) {
    return request({
      url: `${SHOP_BASE_URL}`,
      method: "post",
      data: data,
    });
  }

  static shopUpdate(id: number, data: ShopForm) {
    return request({
      url: `${SHOP_BASE_URL}/${id}`,
      method: "put",
      data: data,
    });
  }

  static shopDel(ids: string) {
    return request({
      url: `${SHOP_BASE_URL}/${ids}`,
      method: "delete",
    });
  }
}

export default ShopAPI;

/**
 * 部门查询参数
 */
export interface ShopQuery extends PageQuery {
  keywords?: string;
  status?: number;
}

export interface ShopVO {
  id?: number;
  name?: string;
  tel?: string;
  notice?: string;
  distributionFee?: number;
  status?: number;
  createdAt?: Date;
}

export interface ShopForm {
  id?: number;
  name?: string;
  tel?: string;
  notice?: string;
  distributionFee?: number;
  status?: number;
}
