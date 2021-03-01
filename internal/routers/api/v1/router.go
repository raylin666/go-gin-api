package api_v1

import (
	api_v1 "gin-api/app/controller/api/v1"
	"github.com/gin-gonic/gin"
)

func Router(routerGroup *gin.RouterGroup)  {
	home(routerGroup)
}

func home(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	router := routerGroup.Group("/home")
	{
		router.GET("", api_v1.HomeIndex)
	}
	return router
}