package utils

import (
	"encoding/json"
	"go-gin-api/internal/constant"
)

type Error struct {
	Code int
	Message string
}

func (e *Error) Error() string {
	err, _ := json.Marshal(e)
	return string(err)
}

func NewError(code int) *Error {
	return &Error{
		Code: code,
		Message: constant.GetStatusText(code),
	}
}
