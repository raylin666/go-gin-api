package http

import (
	"encoding/xml"
	"fmt"
	"gin-api/app/middleware/http"
	"gin-api/internal/constant"
	"gin-api/internal/utils"
	"sync"
	"time"
)

var (
	outputPool *sync.Pool
	// 获取返回格式类型集合
	Format = []string{FormatJSON, FormatXML, FormatYAML, FormatJSONP}
)

const (
	FormatJSON  = "JSON"
	FormatXML   = "XML"
	FormatYAML  = "YAML"
	FormatJSONP = "JSONP"
)

type H map[string]interface{}

// 输出结构体
type Output struct {
	Builder Builder

	// 输出格式, 如 JSON XML JSONP
	Format  string
}

// 构建响应结构体
type Builder struct {
	// HTTP 状态码
	HttpCode int
	// 业务状态码
	Code int
	// 状态码提示信息
	Message string
	// 响应内容信息
	Data H
	// 响应总时长
	ResponseTime time.Duration
}

func init()  {
	outputPool = &sync.Pool{
		New: func() interface {} {
			return new(Output)
		},
	}
}

func (output *Output) reset() {
	output.Format = FormatJSON
	output.Builder.HttpCode = 0
	output.Builder.Code = 0
	output.Builder.Message = ""
	output.Builder.Data = map[string]interface{}{}
	output.Builder.ResponseTime = 0
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
func Response(ctx *utils.Context, out *Output)  {
	code, output := handlerResponse(out)

	output.Builder.ResponseTime = time.Now().Sub(http.GetRequest().StartTime)

	switch output.Format {
	case FormatJSON:
		builderResponseJSON(ctx, code, output)
	case FormatJSONP:
		builderResponseJSONP(ctx, code, output)
	case FormatXML:
		builderResponseXML(ctx, code, output)
	case FormatYAML:
		builderResponseYAML(ctx, code, output)
	default:
		builderResponseJSON(ctx, code, output)
	}

	output.reset()
}

func Set(builder Builder, format string) *Output {
	output := outputPool.Get().(*Output)
	output.Builder = builder
	output.Format = format
	outputPool.Put(output)
	return output
}

func (output *Output) Success(ctx *utils.Context, h H) {
	output.Builder.Data = h
	Response(ctx, output)
}

func (output *Output) Error(ctx *utils.Context, code int) {
	output.Builder.Code = code
	Response(ctx, output)
}

func SuccessResponse(ctx *utils.Context, h H) {
	output := outputPool.Get().(*Output)
	output.Builder = Builder{
		Data: h,
	}
	outputPool.Put(output)
	Response(ctx, output)
}

func ErrorResponse(ctx *utils.Context, code int) {
	output := outputPool.Get().(*Output)
	output.Builder = Builder{
		Code: code,
	}
	outputPool.Put(output)
	Response(ctx, output)
}

func handlerResponse(output *Output) (code int, out *Output) {
	if output.Builder.Code == 0 {
		output.Builder.Code = constant.StatusOK
	}

	// code HTTP 状态码
	code = output.Builder.Code
	if code > 600 {
		if output.Builder.HttpCode == 0 {
			code = constant.StatusOK
		} else {
			code = output.Builder.HttpCode
		}
	}

	if output.Builder.Message == "" {
		output.Builder.Message = constant.GetStatusText(output.Builder.Code)
	}

	return code, output
}

// MarshalXML allows type H to be used with xml.Marshal.
func (h H) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{
		Space: "",
		Local: "map",
	}
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	for key, value := range h {
		elem := xml.StartElement{
			Name: xml.Name{Space: "", Local: key},
			Attr: []xml.Attr{},
		}
		if err := e.EncodeElement(value, elem); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func builderResponseJSON(ctx *utils.Context, code int, out *Output) {
	ctx.JSON(code, builderResponse(
		out.Builder.Code,
		out.Builder.Message,
		out.Builder.Data,
		out.Builder.ResponseTime))
}

func builderResponseXML(ctx *utils.Context, code int, out *Output) {
	ctx.XML(code, builderResponse(
		out.Builder.Code,
		out.Builder.Message,
		out.Builder.Data,
		out.Builder.ResponseTime))
}

func builderResponseYAML(ctx *utils.Context, code int, out *Output) {
	ctx.YAML(code, builderResponse(
		out.Builder.Code,
		out.Builder.Message,
		out.Builder.Data,
		out.Builder.ResponseTime))
}

func builderResponseJSONP(ctx *utils.Context, code int, out *Output) {
	ctx.JSONP(code, builderResponse(
		out.Builder.Code,
		out.Builder.Message,
		out.Builder.Data,
		out.Builder.ResponseTime))
}

// 构建响应数据
func builderResponse(code int, message string, data H, responseTime time.Duration) H {
	return H{
		"code":         code,
		"message":      message,
		"data":         data,
		"responseTime": fmt.Sprintf("%s", responseTime),
	}
}


