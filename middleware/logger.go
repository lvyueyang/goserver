package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"selfserver/lib/logs"
	"strings"
	"time"
)

// 自定义一个结构体，实现 gin.ResponseWriter interface
type responseWriter struct {
	gin.ResponseWriter
	b *bytes.Buffer
}

// 重写 Write([]byte) (int, error) 方法
func (w responseWriter) Write(b []byte) (int, error) {
	//向一个bytes.buffer中写一份数据来为获取body使用
	w.b.Write(b)
	//完成gin.Context.Writer.Write()原有功能
	return w.ResponseWriter.Write(b)
}

type LoggerFormatterParams struct {
	Request      *http.Request
	TimeStamp    time.Time      // 服务器返回响应后的时间戳
	StatusCode   int            // HTTP 响应代码
	Duration     time.Duration  // 表示服务器处理请求所花费的时间
	ClientIP     string         // 客户端的 IP 地址
	Method       string         // 请求的 HTTP 方法
	Path         string         // 客户端请求的路径
	ErrorMessage string         // 在处理请求时发生的错误消息
	ResponseSize int            // 响应体的大小
	Response     string         // 响应体数据
	Query        string         // Query 请求参数
	Body         string         // Body 请求参数
	Keys         map[string]any // 在请求的上下文中设置的键值对
}

func RequestLogger() gin.HandlerFunc {
	return loggerMiddleware()
}

func loggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		bodyJson := body2string(c)

		writer := responseWriter{
			c.Writer,
			bytes.NewBuffer([]byte{}),
		}
		c.Writer = writer

		c.Next()

		param := LoggerFormatterParams{
			Request:      c.Request,
			Keys:         c.Keys,
			ClientIP:     c.ClientIP(),
			Method:       c.Request.Method,
			StatusCode:   c.Writer.Status(),
			ResponseSize: c.Writer.Size(),
			Response:     writer.b.String(),
			Path:         c.Request.URL.Path,
			Query:        c.Request.URL.RawQuery,
			Body:         bodyJson,
		}

		// 结束时间
		param.TimeStamp = time.Now()
		param.Duration = param.TimeStamp.Sub(start)

		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()

		// Path 是 /api 开头的
		if strings.HasPrefix(param.Path, "/api") {
			if param.StatusCode == http.StatusOK {
				fmt.Println("请求成功", param.Path)
				logs.Info().
					Str("path", param.Path).
					Str("method", param.Method).
					Str("ip", param.ClientIP).
					Dur("duration", param.Duration).
					Time("timeStamp", param.TimeStamp).
					Int("status", param.StatusCode).
					Int("responseSize", param.ResponseSize).
					Str("response", param.Response).
					Str("body", param.Body).
					Str("query", param.Query).
					Msg("RequestSuccess")

			} else {
				logs.Err().
					Str("path", param.Path).
					Str("method", param.Method).
					Str("ip", param.ClientIP).
					Dur("duration", param.Duration).
					Time("timeStamp", param.TimeStamp).
					Int("status", param.StatusCode).
					Int("responseSize", param.ResponseSize).
					Str("response", param.Response).
					Str("body", param.Body).
					Str("error", param.ErrorMessage).
					Str("query", param.Query).
					Msg("RequestSuccess")
			}
		}
	}
}

func body2string(c *gin.Context) string {
	var body map[string]any
	err := c.ShouldBindBodyWith(&body, binding.JSON)
	if err != nil {
		return ""
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		return ""
	}
	return string(jsonData)
}
