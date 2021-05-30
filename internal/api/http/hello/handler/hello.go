package handler

import (
	"go-gin-api/internal/api/context"
	"go-gin-api/internal/api/http/hello/logic"
	"go-gin-api/internal/api/http/hello/types/params"
	"go-gin-api/internal/api/utils"
	"go-gin-api/internal/constant"
)

// Hello GO.
func Hello(ctx *context.Context) {
	var req params.HelloReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Error(constant.StatusParamsParseError)
		return
	}
	utils.HandlerResponse(ctx, req, func(req interface{}) (resp interface{}, err *utils.Error) {
		l := logic.NewHelloLogic(ctx)
		resp, err = l.HelloLogic(req.(params.HelloReq))
		return resp, err
	})
}



