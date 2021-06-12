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
	var message string
	res_content := l.ctx.Model.KeywordTests.GetFirst(req.KeyWord)
	if res_content.Keyword == "" {
		message = "没有查询到数据哦 ~"
	} else {
		message = res_content.ResContent
	}

	return &params.HelloResp{
		Message: message,
	}, nil
}
