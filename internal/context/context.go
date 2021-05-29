package context

import (
	"github.com/gin-gonic/gin"
)

// 上下文结构体
type Context struct {
	*gin.Context
}

// 上下文处理函数
type ContextHandlerFunc func(ctx *Context)

// 上下文处理
func ContextHandler(handler ContextHandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var context = new(Context)
		context.Context = ctx
		handler(context)
	}
}
