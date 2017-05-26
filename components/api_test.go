package components

import (
	"net/http"
	"testing"
)

func TestWriteApiResponse(t *testing.T) {
	res := ApiResponse{
		StatusCode: http.StatusOK,
		Message:    "ok",
	}
	responseData := ParseApiResponse(res)
	if responseData["message"].(string) != "ok" {
		t.Error("error in response message", responseData)
	}

	// error response
	jsonParseError := errors.JSONParseError{}
	res = ApiResponse{
		Error: jsonParseError,
	}
	responseData = ParseApiResponse(res)
	if responseData["message"].(string) != jsonParseError.Error() {
		t.Error("error in response message", responseData)
	}
}
