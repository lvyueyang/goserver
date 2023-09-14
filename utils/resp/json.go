package resp

import (
	"net/http"
)

const (
	DefaultSuccessMsg    = "success"
	DefaultErrorMsg      = "success"
	DefaultParamErrorMsg = "请求参数错误"
)

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// Success 成功返回
func Success(data any, msg string) (int, Result) {
	return http.StatusOK, Result{
		Code: 200,
		Msg:  msg,
		Data: data,
	}
}

// Succ 成功返回的简写形式 使用 默认值作为 msg
func Succ(data any) (int, Result) {
	return Success(data, DefaultSuccessMsg)
}

// SuccNil 成功返回的简写形式 返回 空值 和 默认值作为 msg
func SuccNil() (int, Result) {
	return Success(nil, DefaultSuccessMsg)
}

// Err 错误返回
func Err(data any, msg string, code int) Result {
	return Result{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// ParamErr 请求参数错误返回
func ParamErr(msg string) (int, Result) {
	return http.StatusBadRequest, Err(nil, msg, http.StatusBadRequest)
}

// ServerErr 服务器错误返回
func ServerErr(data any, msg string, code int) (int, Result) {
	return http.StatusInternalServerError, Err(data, msg, code)
}

func ParseErr(err error) (int, Result) {
	return ServerErr(err, err.Error(), http.StatusBadRequest)
}
