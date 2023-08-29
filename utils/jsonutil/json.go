package jsonutil

type ResponseData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// SuccessResponse 成功返回
func SuccessResponse(data any, msg string) ResponseData {
	return ResponseData{
		Code: 200,
		Msg:  msg,
		Data: data,
	}
}

// ErrorResponse 错误返回
func ErrorResponse(data any, msg string, code int) ResponseData {
	return ResponseData{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
