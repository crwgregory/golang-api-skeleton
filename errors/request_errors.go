package errors

import (
	"fmt"
	"net/http"
	"strings"
)

type UnauthorizedError struct {
	ApiError
}

func (u UnauthorizedError) GetStatusCode() int {
	return http.StatusUnauthorized
}
func (u UnauthorizedError) Error() string {
	return "You are not authorized to access this resource."
}

type MissingRequiredRequestInformationError struct {
	Needs []string
}

func (m MissingRequiredRequestInformationError) Error() string {
	return fmt.Sprintf("Your request is missing required information: %s", strings.Join(m.Needs, ","))
}
func (m MissingRequiredRequestInformationError) GetStatusCode() int {
	return http.StatusUnprocessableEntity
}
