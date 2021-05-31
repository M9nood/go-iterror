package iterror

import (
	"net/http"
	"strconv"
)

type ErrorException interface {
	Error() string
	GetCode() string
	GetHttpCode() int
	GetName() string
}

type ErrorInfo struct {
	ErrorCode     string
	ErrorHttpCode int
	ErrorName     string
	ErrorMessage  string
}

func ErrorHandler(errCode string) ErrorInfo {
	if _, ok := errors[errCode]; ok {
		return errors[errCode]
	}
	return errors["500"]
}

func NewError(errCode string, msg string) *ErrorInfo {
	err := ErrorHandler(errCode)
	err.ErrorMessage = msg
	return &err
}

func NewErrorCustomHttpCode(errCode string, httpCode int, msg string) *ErrorInfo {
	err := ErrorHandler(errCode)
	err.ErrorMessage = msg
	err.ErrorHttpCode = httpCode
	return &err
}

func (e *ErrorInfo) Error() string {
	if e.ErrorMessage == "" {
		i, _ := strconv.Atoi(e.ErrorCode)
		return http.StatusText(i)
	}
	return e.ErrorMessage
}

func (e *ErrorInfo) GetCode() string {
	return e.ErrorCode
}

func (e *ErrorInfo) GetHttpCode() int {
	return e.ErrorHttpCode
}

func (e *ErrorInfo) GetName() string {
	return e.ErrorName
}
