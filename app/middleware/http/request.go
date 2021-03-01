package http

import (
	"github.com/gin-gonic/gin"
	"time"
)

var (
	request Request
)

type Request struct {
	requestTime
}

type requestTime struct {
	// 请求开始时间
	StartTime time.Time

	// 请求结束时间
	EndTime time.Time

	// 请求总时长
	LatencyTime time.Duration
}

func RequestMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 请求之前
		beforeRequestMiddleware(context)

		context.Next()

		// 请求之后
		afterRequestMiddleware(context)
	}
}

func beforeRequestMiddleware(context *gin.Context) {
	request.StartTime = time.Now()
}

func afterRequestMiddleware(context *gin.Context) {
	// 计算请求时间
	request.EndTime = time.Now()
	request.LatencyTime = request.EndTime.Sub(request.StartTime)
}

// 获取请求信息
func GetRequest() Request {
	return request
}

