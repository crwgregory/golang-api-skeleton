package errors

import "net/http"

type DataConversionError struct {
	ApiError
}

func (c DataConversionError) GetStatusCode() int {
	if c.StatusCode == 0 {
		return http.StatusInternalServerError
	}
	return c.StatusCode
}
func (c DataConversionError) Error() string {
	if c.Message == "" {
		return "Something went wrong trying to convert some data."
	}
	return c.Message
}

type DataNotFoundError struct {
	ApiError
	DataName string
}

func (c DataNotFoundError) GetStatusCode() int {
	return http.StatusNotFound
}
func (c DataNotFoundError) Error() string {
	if c.Message != "" {
		return c.Message
	}
	if c.DataName != "" {
		return c.DataName + " not found."
	}
	return "Data not found."
}
