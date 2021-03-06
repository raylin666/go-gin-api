package service

import (
	"gin-api/internal/http"
	"sync"
)

var servicePool *sync.Pool

type ServiceResponse struct {
	// 是否成功的返回
	OK bool
	// 成功返回的内容
	Data http.H
	// 是否返回的状态码
	Code int
}

func init() {
	servicePool = &sync.Pool{
		New: func() interface{} {
			return new(ServiceResponse)
		},
	}
}

func Success(h http.H) *ServiceResponse {
	response := servicePool.Get().(*ServiceResponse)
	defer servicePool.Put(response)
	response.OK = true
	response.Data = h
	return response
}

func Error(code int) *ServiceResponse {
	response := servicePool.Get().(*ServiceResponse)
	defer servicePool.Put(response)
	response.OK = false
	response.Code = code
	return response
}
