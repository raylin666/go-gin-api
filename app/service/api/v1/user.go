package v1

import (
	userModel "gin-api/app/model/user"
	"gin-api/app/service"
	"gin-api/internal/http"
	"gin-api/internal/utils"
)

func UserInfo(ctx *utils.Context, uid uint64) *service.ServiceResponse {
	var user []userModel.User
	if uid <= 0 {
		user = utils.GetContextUser(ctx)
	} else {
		user = userModel.GetUserOne(uid)
	}

	return service.Success(http.H{
		"user": user,
	})
}