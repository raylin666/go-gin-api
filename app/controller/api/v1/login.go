package v1

import (
	v1 "gin-api/app/service/api/v1"
	"gin-api/app/validator"
	"gin-api/app/validator/request"
	"gin-api/internal/http"
	"gin-api/internal/utils"
)

func LoginIndex(ctx *utils.Context)  {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	ok := validator.Validate(ctx, request.LoginValidate{
		Username: username,
		Password: password,
	})

	if ok {
		response := v1.Login(username, password)
		if response.OK {
			http.SuccessResponse(ctx, response.Data)
		} else {
			http.ErrorResponse(ctx, response.Code)
		}
	}
}