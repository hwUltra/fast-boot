import request from "@/utils/request";
class UserAPI {
  static getList(queryParams?: UserQuery) {
    return request<any, PageResult<UserPageVO[]>>({
      url: "/user",
      method: "get",
      params: queryParams,
    });
  }

  static getUserInfo(userId: number) {
    return request<any, UserForm>({
      url: "/user/" + userId,
      method: "get",
    });
  }

  static addUser(data: UserForm) {
    return request({
      url: "/user",
      method: "post",
      data: data,
    });
  }

  static updateUser(data: UserForm) {
    return request({
      url: "/user",
      method: "put",
      data: data,
    });
  }

  static delUser(ids: string) {
    return request({
      url: "/user/" + ids,
      method: "delete",
    });
  }

  static getThirdList(queryParams: UserQuery) {
    return request<any, PageResult<UserThirdPageVO[]>>({
      url: "/user/third",
      method: "get",
      params: queryParams,
    });
  }
}

export default UserAPI;
/**
 * 登录用户信息
 */
export interface UserInfo {
  userId?: number;
  username?: string;
  nickname?: string;
  avatar?: string;
  idCord?: string;
  gender?: number;
  signature?: string;
  birthday?: string;
  tags?: string;
  source?: string;
  sourceUid?: number;
}

/**
 * 用户查询对象类型
 */
export interface UserQuery extends PageQuery {
  keywords?: string;
}

/**
 * 用户分页对象
 */
export interface UserPageVO {
  /**
   * 用户头像地址
   */
  avatar?: string;
  /**
   * 创建时间
   */
  createAt?: Date;
  /**
   * 性别
   */
  genderLabel?: string;
  /**
   * 用户ID
   */
  id?: number;
  /**
   * 手机号
   */
  mobile?: string;
  /**
   * 用户昵称
   */
  nickname?: string;

  /**
   * 用户名
   */
  username?: string;
}

/**
 * 用户表单类型
 */
export interface UserForm {
  /**
   * 用户ID
   */
  id?: number;
  /**
   * 用户头像
   */
  avatar?: string;
  /**
   * 性别
   */
  gender?: number;
  mobile?: string;
  /**
   * 昵称
   */
  nickname?: string;
  /**
   * 用户名
   */
  username?: string;
}

export interface UserThirdPageVO {
  id?: number;
  uid?: number;
  platform?: string;
  openid?: string;
  unionid?: string;
  nickname?: string;
  avatar?: string;
  createAt?: Date;
}
