package routers

import (
	"gin-api/app/middleware/http"
	"gin-api/internal/env"
	api_v1 "gin-api/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var environment = gin.DebugMode
	switch env.GetEnvironment().Value() {
	case env.ENVIRONMENT_PROD:
		environment = gin.ReleaseMode
		break
	case env.ENVIRONMENT_DEV:
		environment = gin.DebugMode
		break
	case env.ENVIRONMENT_TEST:
		environment = gin.ReleaseMode
		break
	default:
	}

	gin.SetMode(environment)

	engine := gin.New()
	engine.Use(gin.Logger())
	// engine.Use(logger.LoggerToFile())
	engine.Use(http.RequestMiddleware())
	engine.Use(gin.Recovery())

	// API 服务路由
	apiServerRouter(engine)

	return engine
}

func apiServerRouter(engine *gin.Engine)  {
	router := engine.Group("/api")
	{
		// v1
		routerGroup := router.Group("/v1")
		{
			api_v1.Router(routerGroup)
		}
	}
}
