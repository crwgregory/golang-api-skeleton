package components

import (
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/crwgregory/golang-api-skeleton/config"
)

type ApiResponse struct {
	Headers map[string]string
	StatusCode int
	Message string
	Data interface{}
	Error error
	Stack []byte
}

func WriteApiResponse(w http.ResponseWriter, apiResponse ApiResponse) {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(apiResponse.StatusCode)

	for k, v := range apiResponse.Headers {
		fmt.Println(k, v)
		w.Header().Add(k, v)
	}

	response := make(map[string]interface{})

	if apiResponse.Message != "" {
		response["message"] = GetResponseMessage(apiResponse)
	}
	if apiResponse.Data != nil {
		response["data"] = apiResponse.Data
	}

	if apiResponse.Error != nil && config.IsVerbose() {
		response["error"] = apiResponse.Error.Error()
		response["stack"] = string(apiResponse.Stack)
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func GetResponseMessage(res ApiResponse) string {
	if res.Message != "" {
		return res.Message
	}

	switch res.StatusCode {
	case http.StatusOK:
		return "ok"
	}
	// todo: add more

	return "default message"
}
