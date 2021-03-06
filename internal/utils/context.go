package utils

import (
	userModel "gin-api/app/model/user"
	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context

	// 保存用户信息
	User []userModel.User
}

type ContextHandlerFunc func(ctx *Context)

// 设置用户信息
func SetContextUser(ctx *Context, user []userModel.User) {
	ctx.Set("USER", user)
}

// 获取用户信息
func GetContextUser(ctx *Context) []userModel.User {
	return ctx.User
}

// 上下文函数处理
func ContextHandler(handler ContextHandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		Context := new(Context)
		Context.Context = ctx
		if user, ok := ctx.Keys["USER"]; ok {
			Context.User = user.([]userModel.User)
		}
		handler(Context)
	}
}