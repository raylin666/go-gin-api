package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/internal/api/middleware"
	"go-gin-api/internal/constant"
	"go-gin-api/internal/environment"
	hello_router "go-gin-api/internal/api/http/hello/router"
)

type Router struct {}

// 创建路由
func (r *Router) New() *gin.Engine {
	var currentEnvironment = gin.ReleaseMode
	switch environment.GetEnvironment().Value() {
	case constant.EnvironmentPre, constant.EnvironmentProd:
		currentEnvironment = gin.ReleaseMode
	case constant.EnvironmentDev:
		currentEnvironment = gin.DebugMode
	case constant.EnvironmentTest:
		currentEnvironment = gin.TestMode
	default:
	}

	gin.SetMode(currentEnvironment)

	engine := gin.New()

	// 可加载公共中间件配置
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(middleware.RequestLoggerWrite())

	// 可加载路由配置
	loadRouter(engine)

	return engine
}

func loadRouter(engine *gin.Engine) {
	// 注册 Hello 服务接口路由
	hello_router.Router(engine)
}