package api_v1

import (
	api_v1 "gin-api/app/controller/api/v1"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup)  {
	home(r)
}

func home(r *gin.RouterGroup) *gin.RouterGroup {
	router := r.Group("/home")
	{
		router.GET("", api_v1.HomeIndex)
	}
	return router
}