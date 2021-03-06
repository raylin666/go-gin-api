package v1

import (
	v1 "gin-api/app/service/api/v1"
	"gin-api/internal/http"
	"gin-api/internal/utils"
	"strconv"
	"strings"
)

func UserInfo(ctx *utils.Context)  {
	var uid uint64
	uidString := strings.Trim(ctx.Param("uid"), "/")
	if uidString == "" {
		uid = 0
	} else {
		uidInt, _ := strconv.Atoi(uidString)
		uid = uint64(uidInt)
	}

	response := v1.UserInfo(ctx, uid)
	if response.OK {
		http.SuccessResponse(ctx, response.Data)
	} else {
		http.ErrorResponse(ctx, response.Code)
	}
}