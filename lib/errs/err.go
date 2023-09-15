package errs

import "fmt"

// ServerError 服务端错误, 不向用户展示错误详情
type ServerError struct {
	Msg  string
	Err  error
	Info any
}

func (e *ServerError) Error() string {
	return e.Msg
}

func (e *ServerError) Unwrap() error {
	return e.Err
}

// ClientError 客户端错误, 可以向用户展示错误详情
type ClientError struct {
	Msg  string
	Info any
}

func (e *ClientError) Error() string {
	return e.Msg
}

// CreateServerError 创建服务端错误
func CreateServerError(msg string, err error, info any) *ServerError {
	fmt.Printf("ServerError: %+v\n", err)
	fmt.Printf("ServerErrorInfo: %+v\n", info)
	return &ServerError{
		Msg:  msg,
		Err:  err,
		Info: info,
	}
}

// CreateClientError 创建客户端错误
func CreateClientError(msg string, info any) *ClientError {
	return &ClientError{
		Msg:  msg,
		Info: info,
	}
}
