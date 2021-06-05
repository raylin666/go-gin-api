package svc

import "go-gin-api/internal/context"

type Context struct {
	*context.Context
}

func NewContext() *Context {
	return &Context{}
}