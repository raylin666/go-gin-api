package apiv1

import (
	"gin-api/internal/http"
	"github.com/gin-gonic/gin"
)

func HomeIndex(ctx *gin.Context)  {
	http.SuccessResponse(ctx, http.H{
		"message": "hello gin-api.",
	})
}