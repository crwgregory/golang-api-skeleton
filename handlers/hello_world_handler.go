package handlers

import (
	"github.com/crwgregory/golang-api-skeleton/components"
	"net/http"
	"net/url"
)

var hello_world_routes = HandlerCallbackRoutes{
	{
		Name:     "hello",
		Method:   "GET",
		Path:     "/hello/world",
		Callback: helloWorld,
	},
}

type HelloWorldHandler struct {
	Handler
}

// Handle The default Handle function, to be overridden by specific handlers
func (h *HelloWorldHandler) Handle(request Request, logChan chan components.Log) components.ApiResponse {
	return h.RouteController("Hello", hello_world_routes, request, logChan)
}

func helloWorld(r *http.Request, p url.Values) components.ApiResponse {
	return components.ApiResponse{
		Message:    "Hello World!",
		StatusCode: http.StatusOK,
	}
}
