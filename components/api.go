package components

import (
	"github.com/crwgregory/golang-api-skeleton/config"
	"github.com/crwgregory/golang-api-skeleton/errors"
	"net/http"
)

type ApiResponse struct {
	Headers    map[string]string
	StatusCode int
	Message    string
	Data       interface{}
	Error      error
	Stack      []byte
}

func ParseApiResponse(apiResponse ApiResponse) (response map[string]interface{}) {

	response = make(map[string]interface{})

	if mes := GetResponseMessage(apiResponse); mes != "" {
		response["message"] = mes
	}

	if apiResponse.Data != nil {
		response["data"] = apiResponse.Data
	}

	if apiResponse.Error != nil && config.IsVerbose() {
		response["error"] = apiResponse.Error.Error()
		if stack := string(apiResponse.Stack); stack != "" {
			response["stack"] = stack
		}
		e, ok := apiResponse.Error.(errors.ApiErrorInterface)
		if ok && e.GetPrevious() != nil {
			response["previous"] = e.GetPrevious().Error()
		}
	}
	return
}

func GetResponseMessage(res ApiResponse) string {
	if res.Message != "" {
		return res.Message
	}

	switch res.StatusCode {
	case http.StatusOK:
		return "ok"
	case http.StatusForbidden:
		return "forbidden"
	case http.StatusCreated:
		return "created"
	case http.StatusNoContent:
		return "ok no body"
	case http.StatusNotModified:
		return "not modified"
	case http.StatusBadRequest:
		return "bad request"
	case http.StatusUnauthorized:
		return "Unauthorized"
	case http.StatusNotFound:
		return "resource not found"
	case http.StatusMethodNotAllowed:
		return "method not allowed"
	case http.StatusConflict:
		return "conflict"
	case http.StatusUnsupportedMediaType:
		return "unsupported media"
	case http.StatusUnprocessableEntity:
		return "unprocessable entity"
	case http.StatusTooManyRequests:
		return "to many requests"
	case http.StatusInternalServerError:
		return "server error"
	}
	// todo: add more

	// see if we are handling an known error
	if res.Error != nil {
		_, ok := res.Error.(errors.ApiErrorInterface)
		if ok {
			return res.Error.Error()
		}
	}

	return "hi :)"
}
