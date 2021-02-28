package apiv1

import (
	"gin-api/internal/http"
	"github.com/gin-gonic/gin"
)

func HomeIndex(ctx *gin.Context)  {
	http.Response(ctx, http.Output{
		Builder: http.Builder{
			Data: http.H{
				"message": "hello gin-api.",
			},
		},
	})
}