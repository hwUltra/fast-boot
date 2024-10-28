import request from "@/utils/request";
class CategoryAPI {
  static getCategoryList(queryParams?: CategoryQuery) {
    return request<any, PageResult<CategoryVO[]>>({
      url: "/shop/category/list",
      method: "get",
      params: queryParams,
    });
  }
  static getCategoryInfo(id: number) {
    return request<any, CategoryVO>({
      url: "/shop/category/" + id,
      method: "get",
    });
  }

  static categoryAdd(data: CategoryForm) {
    return request({
      url: "/shop/category/add",
      method: "post",
      data: data,
    });
  }

  static categoryUpdate(data: CategoryForm) {
    return request({
      url: "/shop/category/update",
      method: "post",
      data: data,
    });
  }

  static categoryDel(ids: string) {
    return request({
      url: "/shop/category/del/" + ids,
      method: "get",
    });
  }

  static categoryStatus(id: number, status: number) {
    return request({
      url: "/shop/category/status/" + id + "/" + status,
      method: "get",
    });
  }

  static getCategoryOptions(id: number) {
    return request<any, OptionType[]>({
      url: "/shop/category/options/" + id,
      method: "get",
    });
  }
}
export default CategoryAPI;

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
