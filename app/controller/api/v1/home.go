package v1

import (
	"gin-api/internal/http"
	"gin-api/internal/utils"
)

func HomeIndex(ctx *utils.Context)  {
	http.SuccessResponse(ctx, http.H{
		"message": "hello gin-api.",
	})
}