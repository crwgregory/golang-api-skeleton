package errors

import (
	"net/http"
)

type ApiErrorInterface interface {
	GetStatusCode() int
	Error() string
	GetPrevious() error
}

type ApiError struct {
	Previous   error
	StatusCode int
	Message    string
	Stack      []byte
}

func (a ApiError) Error() string {
	return a.Previous.Error()
}
func (a ApiError) GetStatusCode() int {
	return http.StatusInternalServerError
}
func (a ApiError) GetPrevious() error {
	return a.Previous
}

type LogicError struct {
	ApiError
}

func (a LogicError) Error() string {
	if a.Message == "" {
		return "logic error"
	}
	return a.Message
}
func (a LogicError) GetStatusCode() int {
	if a.StatusCode == 0 {
		return http.StatusInternalServerError
	}
	return a.StatusCode
}
