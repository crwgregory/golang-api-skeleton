package handlers

import (
	"net/http"
	"net/url"
	"strings"
	"github.com/crwgregory/golang-api-skeleton/components"
	"fmt"
	"github.com/crwgregory/golang-api-skeleton/connection"
	"time"
)

type Request struct {
	Request  *http.Request
	Response chan components.ApiResponse
	Route    HandlerRoute
	When     time.Time
}

type HandlerInterface interface {
	Handle(Request, chan components.Log) components.ApiResponse
}

type Handler struct {
	Dynamo *connection.DynamoDB
}

// Handle The default Handle function, to be overridden by specific handlers
func (h *Handler) Handle(request Request, logChan chan components.Log) components.ApiResponse {
	return components.ApiResponse{
		StatusCode:200,
		Message:"this is the defualt Handler, Handle func",
	}
}

// HandlerRoute and pat_routes are used to register the controllers/handlers
type HandlerRoute struct {
	Name    string
	Path    string
	Handler HandlerInterface
}

// HandlerRoutes is an array of HandlerRoute
type HandlerRoutes []HandlerRoute

// HandlerCallbackRoute is a specific Route to each controller/handler method
type HandlerCallbackRoute struct {
	Name            string
	Path            string
	Method          string
	Callback        HandlerCallback
}

// HandlerCallbackRoutes is an array of HandlerCallbackRoute
type HandlerCallbackRoutes []HandlerCallbackRoute

// HandlerCallback is what is called to handle the specific request to the Handler
type HandlerCallback func(*http.Request, url.Values) components.ApiResponse

// RouteController routes the incoming request to
func (h *Handler) RouteController(handlerName string, routes HandlerCallbackRoutes, request Request, logChan chan components.Log) components.ApiResponse {

	r := request.Request
	r.ParseForm()
	params := r.Form

	match, _params := findMatch(r.URL.Path, r.Method, routes)

	if match == nil {
		return components.ApiResponse{
			StatusCode: 404,
			Message: fmt.Sprintf("couldn't find match for path: '%s' method: '%s'", r.URL.Path, r.Method),
		}
	}

	for v, p := range _params {
		params[v] = p
	}

	res := match.Callback(r, params)

	log := components.Log{
		HandlerName: handlerName,
		RouteName:match.Name,
		Response:res,
		Request: *r,
		When: request.When,
	}

	logChan <- log // comment this line out to disable request logging
	return res
}

func findMatch(path, method string, checkPaths HandlerCallbackRoutes) (match *HandlerCallbackRoute, params map[string][]string) {

	var potentialMatches HandlerCallbackRoutes

	pathParts := strings.Split(path, "/")[1:]
	pathPartsLen := len(pathParts)

	for _, check := range checkPaths {

		checkParts := strings.Split(check.Path, "/")[1:]
		if len(checkParts) == pathPartsLen {
			potentialMatches = append(potentialMatches, check)
		}
	}

	params = make(map[string][]string)
	var matches []HandlerCallbackRoute

	for _, potentialMatch := range potentialMatches {

		parts := strings.Split(potentialMatch.Path, "/")[1:]

		isMatch := true

		for i, part := range parts {

			if part[0] == ':' {
				// get the name of the part and set it's value

				key := part[1:]
				_, ok := params[key]

				if !ok {
					var values []string
					values = append(values, pathParts[i])
					params[key] = values
				} else {
					params[key] = append(params[key], pathParts[i])
				}
			} else if pathParts[i] != part {
				isMatch = false
				break
			}
		}
		if isMatch == true {
			matches = append(matches, potentialMatch)
		}
	}

	for _, m := range matches {
		if m.Method == method {
			match = &m
			break
		}
	}

	if match == nil {
		return nil, nil
	}

	return
}
