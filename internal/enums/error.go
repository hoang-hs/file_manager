package enums

import (
	"fmt"
	"net/http"
)

type Error interface {
	error
	GetHttpCode() int
	GetMessage() string
}

type CustomError struct {
	HttpCode int
	Message  string
}

func (c CustomError) GetHttpCode() int {
	return c.HttpCode
}

func (c CustomError) GetMessage() string {
	return c.Message
}

func (c CustomError) Error() string {
	return fmt.Sprintf("[%d] %s", c.HttpCode, c.Message)
}

func NewCustomHttpError(httpCode int, message string) *CustomError {
	return &CustomError{
		HttpCode: httpCode,
		Message:  message,
	}
}
func NewCustomSystemError(message string) *CustomError {
	return &CustomError{
		HttpCode: http.StatusInternalServerError,
		Message:  message,
	}
}

var ErrEntityNotFound = NewCustomSystemError("Entity not found")
var ErrSystemError = NewCustomHttpError(http.StatusInternalServerError, "System error")
var ErrUnAuthenticated = NewCustomHttpError(http.StatusUnauthorized, "Check your account")
