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

export interface ShopOptions {
  list?: OptionType[];
}
