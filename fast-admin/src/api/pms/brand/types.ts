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
