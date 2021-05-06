package context

import (
	"github.com/gin-gonic/gin"
	"time"
)

var context = new(Context)

// 上下文结构体
type Context struct {
	*gin.Context

	// 认证 Key
	Authorization string
}

// 上下文处理函数
type ContextHandlerFunc func(ctx *Context)

// 响应数据包内容
type H map[string]interface{}

// 构建响应结构体
type ResponseBuilder struct {
	// HTTP 状态码
	HttpCode int
	// 业务状态码
	Code int
	// 状态码提示信息
	Message string
	// 响应内容信息
	Data H
	// 响应总时长
	ResponseTime time.Duration
}

// 设置 认证 Key
func (ctx *Context) SetContextAuthorization(value string) {
	ctx.Set("Authorization", value)
}

// 获取 认证 Key
func (ctx *Context) GetContextAuthorization() string {
	return ctx.Authorization
}

// 上下文处理
func ContextHandler(handler ContextHandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		context.Context = ctx
		if authorization, ok := ctx.Keys["Authorization"]; ok {
			context.Authorization = authorization.(string)
		}
		handler(context)
	}
}

func (ctx *Context) Response()  {

}
