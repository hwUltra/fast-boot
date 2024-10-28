/**
 * 响应码枚举
 */
export const enum ResultEnum {
  /**
   * 成功
   */
  SUCCESS = 200,
  /**
   * 错误
   */
  ERROR = "B0001",

  /**
   * 令牌无效或过期
   */
  TOKEN_INVALID = 401,
}
