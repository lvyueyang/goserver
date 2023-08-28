package utils

import "github.com/gin-gonic/gin"

type ResponseData struct {
	code int
	msg  string
	data any
}

// SuccessResponse 成功返回
func SuccessResponse(data any, msg string) any {
	return gin.H{"code": 1000, "msg": msg, "data": data}
}
