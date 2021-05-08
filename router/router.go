package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raylin666/go-gin-api/constant"
	"github.com/raylin666/go-gin-api/environment"
)

type Router struct {
	Before func(*gin.Engine)
	After  func(*gin.Engine)
}

// 创建路由
func (r *Router) New() *gin.Engine {
	var currentEnvironment = gin.ReleaseMode
	switch environment.GetEnvironment().Value() {
	case constant.ENVIRONMENT_PRE, constant.ENVIRONMENT_PROD:
		currentEnvironment = gin.ReleaseMode
	case constant.ENVIRONMENT_DEV:
		currentEnvironment = gin.DebugMode
	case constant.ENVIRONMENT_TEST:
		currentEnvironment = gin.TestMode
	default:
	}

	gin.SetMode(currentEnvironment)

	engine := gin.New()

	// 可加载中间件配置
	if r.Before != nil {
		r.Before(engine)
	}

	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	// 可加载路由配置
	if r.After != nil {
		r.After(engine)
	}

	return engine
}
