package logger

import (
	"fmt"
	"gin-api/internal/constant"
	"gin-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"time"
)

// 日志记录
func LoggerWrite() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		logger.NewWrite(constant.LOG_REQUEST).WithFields(logger.Fields{
			"status_code"  : statusCode,
			"latency_time" : fmt.Sprintf("%s", latencyTime),
			"client_ip"    : clientIP,
			"req_method"   : reqMethod,
			"req_uri"      : reqUri,
		}.Fields()).Info()
	}
}
