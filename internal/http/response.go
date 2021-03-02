package http

import (
	"encoding/xml"
	"fmt"
	"gin-api/app/middleware/http"
	"gin-api/internal/constant"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	output Output
	isSetOutput = false
	responseFormat = DefaultResponseFormat
	format = []string{FormatJSON, FormatXML, FormatYAML, FormatJSONP}
)

const (
	FormatJSON  = "JSON"
	FormatXML   = "XML"
	FormatYAML  = "YAML"
	FormatJSONP = "JSONP"

	DefaultResponseFormat = FormatJSON
)

// 获取返回格式类型
func GetFormat() []string {
	return format
}

type Output struct {
	Builder Builder

	// 输出类型, 例如：JSON、XML、YAML
	Format string
}

type H map[string]interface{}

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

// 构建输出数据结构体
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
func Response(ctx *gin.Context, o Output) {
	httpCode, out := handlerOutputResponse(o)

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

func handlerOutputResponse(out Output) (httpCode int, o Output) {
	if out.Format != "" {
		responseFormat = out.Format
	}

	if out.Builder.Code == 0 {
		out.Builder.Code = constant.StatusOK
	}

	httpCode = out.Builder.Code
	if httpCode > 600 {
		if out.Builder.HttpCode == 0 {
			httpCode = constant.StatusOK
		} else {
			httpCode = out.Builder.HttpCode
		}
	}

	if out.Builder.Message == "" {
		out.Builder.Message = constant.GetStatusText(out.Builder.Code)
	}

	return httpCode, out
}

func SetOutputResponse(out Output) {
	_, output = handlerOutputResponse(out)
	isSetOutput = true
}

func GetOutputResponse() *Output {
	return &output
}

func IsSetOutputResponse() bool {
	return isSetOutput
}

/**
	http.SetOutputResponse(http.Output{Builder: http.Builder{Code: 200}, Format: http.FormatJSON})
	http.SuccessResponse(ctx, http.H{
		"message": "hello gin-api.",
	})
 */
func SuccessResponse(ctx *gin.Context, h H)  {
	var out Output
	if IsSetOutputResponse() {
		out = Output{
			Builder: Builder{
				Code: output.Builder.Code,
				Data: h,
			},
			Format: output.Format,
		}
	} else {
		out = Output{
			Builder: Builder{
				Data: h,
			},
		}
	}

	Response(ctx, out)
}

func ErrorResponse(ctx *gin.Context, code int)  {
	var out Output
	if IsSetOutputResponse() {
		out = Output{
			Builder: Builder{
				Code: code,
				Data: output.Builder.Data,
			},
			Format: output.Format,
		}
	} else {
		out = Output{
			Builder: Builder{
				Code: code,
			},
		}
	}

	Response(ctx, out)
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
func builderResponse(code int, message string, data H, responseTime time.Duration) *H {
	return &H{
		"code":         code,
		"message":      message,
		"data":         data,
		"responseTime": fmt.Sprintf("%s", responseTime),
	}
}
