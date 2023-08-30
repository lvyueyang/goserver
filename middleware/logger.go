package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
	"net/http"
	"selfserver/lib/logger"
	"strings"
	"time"
)

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
	Query        string         // Query 请求参数
	Body         string         // Body 请求参数
	Keys         map[string]any // 在请求的上下文中设置的键值对
}

func Logger() gin.HandlerFunc {
	return loggerMiddleware()
}

func loggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		var bodyJson = body2string(c)

		c.Next()

		param := LoggerFormatterParams{
			Request:      c.Request,
			Keys:         c.Keys,
			ClientIP:     c.ClientIP(),
			Method:       c.Request.Method,
			StatusCode:   c.Writer.Status(),
			ResponseSize: c.Writer.Size(),
			Path:         c.Request.URL.Path,
			Query:        c.Request.URL.RawQuery,
			Body:         bodyJson,
		}

		// 结束时间
		param.TimeStamp = time.Now()
		param.Duration = param.TimeStamp.Sub(start)

		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()

		// Path 是 /api 开头的
		if strings.HasPrefix(param.Path, "/api") && param.Method == "POST" {
			if param.StatusCode == http.StatusOK {
				logger.Logger.Infow("RequestSuccess",
					zap.String("path", param.Path),
					zap.String("method", param.Method),
					zap.String("ip", param.ClientIP),
					zap.Duration("duration", param.Duration),
					zap.Time("timeStamp", param.TimeStamp),
					zap.Int("status", param.StatusCode),
					zap.Int("responseSize", param.ResponseSize),
					zap.String("body", param.Body),
					zap.String("query", param.Query),
				)
			} else {
				logger.Logger.Errorw("RequestFail",
					zap.String("path", param.Path),
					zap.String("method", param.Method),
					zap.String("ip", param.ClientIP),
					zap.Duration("duration", param.Duration),
					zap.Time("timeStamp", param.TimeStamp),
					zap.Int("status", param.StatusCode),
					zap.Int("responseSize", param.ResponseSize),
					zap.String("body", param.Body),
					zap.String("query", param.Query),
					zap.String("error", param.ErrorMessage),
				)
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
