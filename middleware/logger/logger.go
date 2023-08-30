package logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type LogFormatterParams struct {
	Request      *http.Request
	TimeStamp    time.Time      // 服务器返回响应后的时间戳
	StatusCode   int            // HTTP 响应代码
	Duration     time.Duration  // 表示服务器处理请求所花费的时间
	ClientIP     string         // 客户端的 IP 地址
	Method       string         // 请求的 HTTP 方法
	Path         string         // 客户端请求的路径
	ErrorMessage string         // 在处理请求时发生的错误消息
	BodySize     int            // 响应体的大小
	Query        string         // Query 请求参数
	Body         string         // Body 请求参数
	Keys         map[string]any // 在请求的上下文中设置的键值对
}

func Logger() gin.HandlerFunc {
	logger, _ := zap.NewDevelopment()

	return loggerMiddleware(logger)
}

func loggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()

		c.Next()

		param := LogFormatterParams{
			Request:    c.Request,
			Keys:       c.Keys,
			ClientIP:   c.ClientIP(),
			Method:     c.Request.Method,
			StatusCode: c.Writer.Status(),
			BodySize:   c.Writer.Size(),
			Path:       c.Request.URL.Path,
			Query:      c.Request.URL.RawQuery,
		}

		// 结束时间
		param.TimeStamp = time.Now()
		param.Duration = param.TimeStamp.Sub(start)

		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()

		logger.Info("Request Logger",
			zap.String("path", param.Path),
			zap.String("method", param.Method),
			zap.String("ip", param.ClientIP),
			zap.Duration("duration", param.Duration),
			zap.Time("timeStamp", param.TimeStamp),
			zap.Int("status", param.StatusCode),
			zap.String("error", param.ErrorMessage),
			zap.Int("bodySize", param.BodySize),
		)
	}
}
