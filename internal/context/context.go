package context

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/raylin666/go-gin-api/internal/consts"
	"github.com/raylin666/go-gin-api/internal/constant"
)

// 上下文结构体
type Context struct {
	*gin.Context

	// 认证 Key
	Authorization string

	// 响应数据包内容
	ResponseBuilder ResponseBuilder
}

// 上下文处理函数
type ContextHandlerFunc func(ctx *Context)

// 设置 认证 Key
func (ctx *Context) SetContextAuthorization(value string) {
	ctx.Set("Authorization", value)
}

// 获取 认证 Key
func (ctx *Context) GetContextAuthorization() string {
	return ctx.Authorization
}

// 上下文处理
func ContextHandler(handler ContextHandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var context = new(Context)
		context.Context = ctx
		if authorization, ok := ctx.Keys["Authorization"]; ok {
			context.Authorization = authorization.(string)
		}
		handler(context)
	}
}

// API 响应信息
func (ctx *Context) Response()  {
	// 处理响应数据包内容
	ctx.handlerResponse()
}

// 处理响应数据包内容
func (ctx *Context) handlerResponse() {
	// 处理响应状态码
	ctx.ResponseBuilder.WithCode(ctx.ResponseBuilder.Code)
	if ctx.ResponseBuilder.Code <= 600 && ctx.ResponseBuilder.Code > 0 {
		ctx.ResponseBuilder.WithHttpCode(ctx.ResponseBuilder.Code)
	} else if ctx.ResponseBuilder.Code > 600 {
		ctx.ResponseBuilder.WithHttpCode(constant.StatusOK)
	} else {
		ctx.ResponseBuilder.WithHttpCode(ctx.ResponseBuilder.HttpCode)
	}

	ctx.ResponseBuilder.WithMessage(ctx.ResponseBuilder.Message)

	ctx.ResponseBuilder.Data = H{
		"code":         ctx.ResponseBuilder.Code,
		"message":      ctx.ResponseBuilder.Message,
		"data":         ctx.ResponseBuilder.Data,
		"responseTime": fmt.Sprintf("%s", ctx.ResponseBuilder.ResponseTime),
	}

	switch ctx.ResponseBuilder.Format {
	case consts.FORMAT_XML:
		ctx.builderResponseXML()
	case consts.FORMAT_YAML:
		ctx.builderResponseYAML()
	case consts.FORMAT_JSONP:
		ctx.builderResponseJSONP()
	default:
		ctx.builderResponseJSON()
	}
}

func (ctx *Context) builderResponseJSON() {
	ctx.Context.JSON(ctx.ResponseBuilder.Code, ctx.ResponseBuilder.Data)
}

func (ctx *Context) builderResponseXML() {
	ctx.Context.XML(ctx.ResponseBuilder.Code, ctx.ResponseBuilder.Data)
}

func (ctx *Context) builderResponseYAML() {
	ctx.Context.YAML(ctx.ResponseBuilder.Code, ctx.ResponseBuilder.Data)
}

func (ctx *Context) builderResponseJSONP() {
	ctx.Context.JSONP(ctx.ResponseBuilder.Code, ctx.ResponseBuilder.Data)
}

