package context

import (
	"encoding/xml"
	"github.com/raylin666/go-gin-api/internal/consts"
	"github.com/raylin666/go-gin-api/internal/constant"
	"time"
)

// 响应数据包内容
type H map[string]interface{}

// 构建响应结构体
type ResponseBuilder struct {
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
	// 响应格式
	Format string
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

func (response *ResponseBuilder) WithHttpCode(code int) {
	if response.HttpCode == 0 {
		response.HttpCode = constant.StatusOK
	} else {
		response.HttpCode = code
	}
}

func (response *ResponseBuilder) GetHttpCode() int {
	return response.HttpCode
}

func (response *ResponseBuilder) WithCode(code int) {
	if response.Code == 0 {
		response.Code = constant.StatusOK
	} else {
		response.Code = code
	}
}

func (response *ResponseBuilder) GetCode() int {
	return response.Code
}

func (response *ResponseBuilder) WithMessage(message string) {
	if response.Message == "" {
		response.Message = constant.GetStatusText(response.Code)
	} else {
		response.Message = message
	}
}

func (response *ResponseBuilder) GetMessage() string {
	return response.Message
}

func (response *ResponseBuilder) WithData(data H) {
	response.Data = data
}

func (response *ResponseBuilder) GetData() H {
	return response.Data
}

func (response *ResponseBuilder) WithResponseTime(duration time.Duration) {
	response.ResponseTime = duration
}

func (response *ResponseBuilder) GetResponseTime() time.Duration {
	return response.ResponseTime
}

func (response *ResponseBuilder) WithFormat(format string) {
	switch format {
	case consts.FORMAT_JSON:
	case consts.FORMAT_JSONP:
	case consts.FORMAT_XML:
	case consts.FORMAT_YAML:
	default:
		format = consts.FORMAT_JSON
	}
	response.Format = format
}

func (response *ResponseBuilder) GetFormat() string {
	return response.Format
}
