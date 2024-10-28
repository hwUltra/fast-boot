package xerr

import (
	"github.com/hwUltra/fb-tools/result"
)

func NewErrCodeMsg(errCode uint32, errMsg string) *result.CodeError {
	return result.NewErrCodeMsg(errCode, errMsg)
}
func NewErrCode(errCode uint32) *result.CodeError {
	return result.NewErrCodeMsg(errCode, MapErrMsg(errCode))
}

func NewErrMsg(errMsg string) *result.CodeError {
	return result.NewErrCodeMsg(SERVER_COMMON_ERROR, errMsg)
}
