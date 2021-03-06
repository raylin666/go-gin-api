package v1

import (
	"gin-api/app/controller/api/v1"
	"gin-api/app/middleware/jwt"
	"gin-api/internal/utils"
	"github.com/gin-gonic/gin"
)

func Router(routerGroup *gin.RouterGroup) {
	login(routerGroup)
	home(routerGroup)
	user(routerGroup)
}

func login(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	router := routerGroup.Group("/login")
	{
		router.POST("", utils.ContextHandler(v1.LoginIndex))
	}
	return router
}

func home(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	router := routerGroup.Group("/home")
	router.Use(utils.ContextHandler(jwt.JWT()))
	{
		router.GET("", utils.ContextHandler(v1.HomeIndex))
	}
	return router
}

func user(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	router := routerGroup.Group("/user")
	router.Use(utils.ContextHandler(jwt.JWT()))
	{
		router.GET("/info/*uid", utils.ContextHandler(v1.UserInfo))
	}
	return router
}
