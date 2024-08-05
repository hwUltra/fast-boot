export interface CategoryQuery extends PageQuery {
  shopId?: number;
  keywords?: string;
  status?: number;
}

export interface CategoryVO {
  id?: number;
  shopId?: number;
  parentId?: number;
  name?: string;
  level?: number;
  icon?: string;
  visible?: number;
  sort?: number;
  createdAt?: Date;
  children?: CategoryVO[];
}

export interface CategoryForm {
  id?: number;
  shopId?: number;
  parentId?: number;
  name?: string;
  level?: number;
  icon?: string;
  visible?: number;
  sort?: number;
}
