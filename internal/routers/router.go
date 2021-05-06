package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/raylin666/go-gin-api/environment"
	"github.com/raylin666/go-gin-api/internal/consts"
)

// 创建路由
func New() *gin.Engine {
	var currentEnvironment = gin.ReleaseMode
	switch environment.GetEnvironment().Value() {
	case consts.ENVIRONMENT_PRE, consts.ENVIRONMENT_PROD:
		currentEnvironment = gin.ReleaseMode
	case consts.ENVIRONMENT_DEV:
		currentEnvironment = gin.DebugMode
	case consts.ENVIRONMENT_TEST:
		currentEnvironment = gin.TestMode
	default:
	}

	gin.SetMode(currentEnvironment)

	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	return engine
}
