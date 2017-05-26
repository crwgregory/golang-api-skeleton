package errors

import (
	"github.com/crwgregory/golang-api-skeleton/config"
	"net/http"
)

type DatabaseQueryError struct {
	ApiError
	Query        string
	CauseMessage string
}

func (d DatabaseQueryError) GetStatusCode() int {
	return http.StatusInternalServerError
}
func (d *DatabaseQueryError) Error() string {
	e := "There was an error querying the database. " + d.CauseMessage + " "
	if config.IsVerbose() {
		return e + d.Query
	}
	return e
}

type DatabaseUpdateError struct {
	ApiError
	Query        string
	CauseMessage string
}

func (d DatabaseUpdateError) GetStatusCode() int {
	return http.StatusInternalServerError
}
func (d *DatabaseUpdateError) Error() string {
	e := "There was an error updating the database. " + d.CauseMessage + " "
	if config.IsVerbose() {
		return e + d.Query
	}
	return e
}
