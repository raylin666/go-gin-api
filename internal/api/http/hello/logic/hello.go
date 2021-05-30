package logic

import (
	"go-gin-api/internal/api/context"
	"go-gin-api/internal/api/http/hello/types/params"
	"go-gin-api/internal/api/utils"
)

type HelloLogic struct {
	ctx *context.Context
}

func NewHelloLogic(ctx *context.Context) *HelloLogic {
	return &HelloLogic{
		ctx: ctx,
	}
}

// Hello GO.
func (l *HelloLogic) HelloLogic(req params.HelloReq) (*params.HelloResp, *utils.Error) {
	return &params.HelloResp{
		Message: req.KeyWord + ", Hello GO.",
	}, nil
}
