package svc

import (
	"go-gin-api/internal/context"
	"go-gin-api/internal/model"
)

type Context struct {
	*context.Context

	Model struct {
		KeywordTest *model.KeywordTestsModel
	}
}

func NewContext() *Context {
	var ctx = new(Context)
	ctx.Model.KeywordTest = model.NewKeywordTestsModel()
	return ctx
}