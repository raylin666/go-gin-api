package http

import (
	"fmt"
	"gin-api/app/middleware/http"
	"gin-api/internal/constant"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	FormatJSON  = "JSON"
	FormatXML   = "XML"
	FormatYAML  = "YAML"
	FormatJSONP = "JSONP"

	DefaultResponseFormat = FormatJSON
)

var responseFormat = DefaultResponseFormat

var format = []string{FormatJSON, FormatXML, FormatYAML, FormatJSONP}

// 获取返回格式类型
func GetResponseFormat() []string {
	return format
}

type Output struct {
	Builder Builder

	// 输出类型, 例如：JSON、XML、YAML
	Format string
}

type H map[string]interface{}

// 构建输出数据结构体
type Builder struct {
	// 业务状态码
	Code int

	// 状态码提示信息
	Message string

	// 响应内容信息
	Data H

	// 响应总时长
	ResponseTime time.Duration
}

// 响应数据
/**
	http.Response(ctx, http.Output{
		Builder: http.Builder{
			Data: http.H{
				"message": "hello gin-api.",
			},
		},
	})
 */
func Response(ctx *gin.Context, out Output) {
	if out.Format != "" {
		responseFormat = out.Format
	}

	if out.Builder.Code == 0 {
		out.Builder.Code = constant.StatusOK
	}

	httpCode := out.Builder.Code
	if httpCode > 600 {
		httpCode = constant.StatusOK
	}

	if out.Builder.Message == "" {
		out.Builder.Message = constant.GetStatusText(out.Builder.Code)
	}

	out.Builder.ResponseTime = time.Now().Sub(http.GetRequest().StartTime)

	switch responseFormat {
	case FormatJSON:
		builderResponseJSON(ctx, httpCode, out)
	case FormatJSONP:
		builderResponseJSONP(ctx, httpCode, out)
	case FormatXML:
		builderResponseXML(ctx, httpCode, out)
	case FormatYAML:
		builderResponseYAML(ctx, httpCode, out)
	default:
		builderResponseJSON(ctx, httpCode, out)
	}
}

func builderResponseJSON(ctx *gin.Context, httpCode int, out Output) {
	ctx.JSON(httpCode, builderResponse(
		out.Builder.Code,
		out.Builder.Message,
		out.Builder.Data,
		out.Builder.ResponseTime))
}

func builderResponseXML(ctx *gin.Context, httpCode int, out Output) {
	ctx.XML(httpCode, builderResponse(
		out.Builder.Code,
		out.Builder.Message,
		out.Builder.Data,
		out.Builder.ResponseTime))
}

func builderResponseYAML(ctx *gin.Context, httpCode int, out Output) {
	ctx.YAML(httpCode, builderResponse(
		out.Builder.Code,
		out.Builder.Message,
		out.Builder.Data,
		out.Builder.ResponseTime))
}

func builderResponseJSONP(ctx *gin.Context, httpCode int, out Output) {
	ctx.JSONP(httpCode, builderResponse(
		out.Builder.Code,
		out.Builder.Message,
		out.Builder.Data,
		out.Builder.ResponseTime))
}

// 构建响应数据
func builderResponse(code int, message string, data H, responseTime time.Duration) H {
	return H(
		gin.H{
			"code":         code,
			"message":      message,
			"data":         data,
			"responseTime": fmt.Sprintf("%s", responseTime),
		})
}
