package jsonutil

type ResponseData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// SuccessResponse 成功返回
func SuccessResponse(data any, msg string) any {
	return ResponseData{
		Code: 200,
		Msg:  msg,
		Data: data,
	}
}
