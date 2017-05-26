package errors

import (
	"github.com/crwgregory/golang-api-skeleton/config"
	"net/http"
)

type WrongPasswordError struct {
	ApiError
}

func (w WrongPasswordError) Error() string {
	return "Wrong Password"
}
func (w WrongPasswordError) GetStatusCode() int {
	return http.StatusForbidden
}

type UserNotFound struct {
	ApiError
	Query string
}

func (u UserNotFound) Error() string {
	if config.IsVerbose() {
		return "User not found. " + u.Query
	}
	return "User not found."
}
func (u UserNotFound) GetStatusCode() int {
	return http.StatusUnprocessableEntity
}

type PasswordToShort struct {
	ApiError
}

func (p PasswordToShort) Error() string {
	return "Password is to short. Must be at least 8 characters."
}
func (p PasswordToShort) GetStatusCode() int {
	return http.StatusUnprocessableEntity
}

type UserDoesNotHaveSufficientPermissions struct {
	ApiError
}

func (u UserDoesNotHaveSufficientPermissions) Error() string {
	return "You don't have sufficient permissions to access this."
}
func (u UserDoesNotHaveSufficientPermissions) GetStatusCode() int {
	return http.StatusForbidden
}
