/**
 * 响应码枚举
 */
export const enum ResultCodeEnum {
  /**
   * 成功
   */
  SUCCESS = 200,
  /**
   * 错误
   */
  ERROR = 10001,

  /**
   * 令牌无效或过期
   */
  TOKEN_INVALID = 401,
}
