package hello

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/internal/api/http/hello/handler"
	"go-gin-api/internal/api/context"
	"go-gin-api/internal/api/http/hello/middleware"
)

// Hello 服务接口路由
func Router(engine *gin.Engine) {
	// 中间件注册
	engine.Use(middleware.Hello())

	// v1 版本
	router := engine.Group("/hello/v1")
	{
		HelloV1(router)
	}
}

func HelloV1(router *gin.RouterGroup) *gin.RouterGroup {
	{
		router.GET("", context.ContextHandler(handler.Hello))
	}
	return router
}