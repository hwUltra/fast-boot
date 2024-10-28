import request from "@/utils/request";
class BrandAPI {
  static getBrandList(queryParams?: BrandQuery) {
    return request<any, BrandPageResult>({
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
  static getBrandOptions(id: number) {
    return request<any, OptionType[]>({
      url: "/shop/brand/options/" + id,
      method: "get",
    });
  }

  /**
   * 添加品牌
   *
   * @param data
   */
  static brandAdd(data: BrandForm) {
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
  static brandUpdate(data: BrandForm) {
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
  static brandDel(ids: string) {
    return request({
      url: "/shop/brand/del/" + ids,
      method: "get",
    });
  }
}
export default BrandAPI;
/**
 * 品牌查询参数类型声明
 */
export interface BrandQuery extends PageQuery {
  shopId: number;
  keywords?: string;
}

/**
 * 品牌分页列表项声明
 */
export interface BrandInfo {
  id: string;
  shopId: number;
  name: string;
  logo: string;
  sort: number;
}

/**
 * 品牌分页项类型声明
 */
export type BrandPageResult = PageResult<BrandInfo[]>;

/**
 * 品牌表单类型声明
 */
export interface BrandForm {
  id: number | undefined;
  shopId: number;
  name: string;
  logo: string;
  sort: number;
}
