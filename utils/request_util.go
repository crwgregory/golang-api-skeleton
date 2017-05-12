package utils

import (
	"net/url"
	"net/http"
	"encoding/json"
	"io"
	"io/ioutil"
	"bytes"
	"github.com/crwgregory/golang-api-skeleton/components"
	"fmt"
	"strings"
	"runtime/debug"
)

func GetRequestData(keys []string, r *http.Request) (data map[string]interface{}, errRes *components.ApiResponse) {
	needs := RequestHasParams(keys, r)
	if len(needs) > 0 {
		return nil, &components.ApiResponse{
			StatusCode: http.StatusUnprocessableEntity,
			Message: fmt.Sprintf("request requires: %s", strings.Join(needs, ",")),
		}
	}
	data, err := GetJsonBody(r)
	if err != nil {
		return nil, &components.ApiResponse{
			StatusCode:http.StatusInternalServerError,
			Message:"there was an error parsing the request body",
			Error: err,
			Stack: debug.Stack(),
		}
	}
	return
}

func QueryHasParams(requiredParams []string, params url.Values) []string {
	var needs []string
	for _, p := range requiredParams {
		_, ok := params[p]
		if !ok {
			needs = append(needs, p)
		}
	}
	return needs
}

func RequestHasParams(requiredParams []string, r *http.Request) []string {

	var needs []string

	data, err := GetJsonBody(r)
	if err != nil {
		panic(err)
	}

	for _, req := range requiredParams {
		_, ok := data[req]
		if !ok {
			needs = append(needs, req)
		}
	}

	return needs
}

func GetJsonBody(r *http.Request) (map[string]interface{}, error) {
	data := make(map[string]interface{})

	body, err := copyRequestBody(r)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(body)
	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func copyRequestBody(r *http.Request) (copiedBody io.Reader, err error) {
	buf, err := ioutil.ReadAll(r.Body)
	copiedBody = ioutil.NopCloser(bytes.NewBuffer(buf))
	r.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
	return
}
