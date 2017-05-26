package errors

import (
	"fmt"
	"net/http"
)

type JWTTokenExpiredError struct{}

func (j JWTTokenExpiredError) Error() string {
	return "JWT Expired"
}

type JWTSignatureMismatch struct{}

func (j JWTSignatureMismatch) Error() string {
	return "The JWT Signatures Do Not Match"
}

type JSONParseError struct {
	ApiError
}

func (j JSONParseError) Error() string {
	var pre string
	if j.Previous != nil {
		pre = j.Previous.Error()
	}
	return fmt.Sprintf("There was an error parsing the Request to JSON: %s", pre)
}
func (j JSONParseError) GetStatusCode() int {
	return http.StatusUnprocessableEntity
}

type MissingAuthHeader struct {
	ApiError
}

func (m MissingAuthHeader) GetStatusCode() int {
	return http.StatusForbidden
}
func (m MissingAuthHeader) Error() string {
	return "Missing Authorization Header"
}
