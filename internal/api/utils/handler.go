package utils

import (
	"go-gin-api/internal/api/context"
)

// 处理响应
func HandlerResponse(ctx *context.Context, req interface{}, callback func(req interface{}) (resp interface{}, err *Error)) {
	if callback != nil {
		if valid := ctx.RequestValidate(req); valid {
			resp, err := callback(req)
			if err != nil {
				if (err.Code >= 200 && err.Code < 300) || (err.Code == 0) {
					ctx.Success(resp)
					return
				}
				ctx.Error(err.Code)
			} else {
				ctx.Success(resp)
			}
		}
	}
	return
}
