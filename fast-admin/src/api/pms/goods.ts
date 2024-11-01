import request from "@/utils/request";
class GoodsAPI {
  /**
   * 获取商品分页列表
   *
   * @param queryParams
   */
  static getGoodsList(queryParams: GoodsQuery) {
    return request<any, GoodsPageResult>({
      url: "/goods",
      method: "get",
      params: queryParams,
    });
  }

  /**
   * 获取商品详情
   *
   * @param id
   */
  static getGoodsInfo(id: string) {
    return request<any, GoodsDetail>({
      url: "/goods/" + id,
      method: "get",
    });
  }

  /**
   * 添加商品
   *
   * @param data
   */
  static editGoods(data: object) {
    return request({
      url: "/goods/add",
      method: "post",
      data: data,
    });
  }

  /**
   * 删除商品
   *
   * @param ids
   */
  static delGoods(ids: string) {
    return request({
      url: "/goods/del/" + ids,
      method: "get",
    });
  }
}

export default GoodsAPI;

/**
 * 商品查询参数类型声明
 */
export interface GoodsQuery extends PageQuery {
  shopId?: number;
  keywords?: string;
  categoryId?: number;
}

/**
 * 商品列表项类型声明
 */
export interface GoodsInfo {
  id: string;
  name: string;
  categoryId?: any;
  brandId?: any;
  originPrice: string;
  price: string;
  sales: number;
  picUrl?: any;
  album?: any;
  unit?: any;
  description: string;
  detail: string;
  status?: any;
  categoryName: string;
  brandName: string;
  skuList: Sku[];
}

/**
 * 商品规格项类型声明
 */
export interface Sku {
  id: string;
  skuSn?: any;
  name: string;
  spuId?: any;
  specIds: string;
  price: string;
  stockNum: number;
  lockedStockNum?: any;
  picUrl?: any;
}

/**
 * 商品分页项类型声明
 */
export type GoodsPageResult = PageResult<GoodsInfo[]>;

/**
 * 商品表单数据类型声明
 */
export interface GoodsDetail {
  id?: number;
  shopId?: number;
  name?: string;
  categoryId?: string;
  brandId?: string;
  originPrice?: number;
  price?: number;
  picUrl?: string;
  album: string[];
  description?: string;
  detail?: string;
  attrList: any[];
  specList: any[];
  skuList: any[];
}
