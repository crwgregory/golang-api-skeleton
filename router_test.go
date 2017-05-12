package main

import (
	"testing"
	"github.com/crwgregory/golang-api-skeleton/handlers"
	"github.com/crwgregory/golang-api-skeleton/components"
	"net/http"
	"net/url"
)

func TestBuildRouter(t *testing.T) {
	r := BuildRouter(handlers.HandlerRoutes{
		handlers.HandlerRoute{
			"Test",
			"/test",
			new(handlers.TestHandler),
		},
	})

	testUrl, err := url.Parse("http://localhost:8080/test")

	if err != nil {
		panic(err)
	}

	req := &http.Request{
		Method: "GET",
		URL: testUrl,
	}

	limit := 100

	response := make(chan components.ApiResponse)

	// put the request into the Routers request channel with the created response channel for this request
	for i := 0; i < limit; i++ {
		r.request <- handlers.Request{
			Request: req,
			Response: response,
		}
	}

	count := 0

	for range response {
		count++
		if count == limit -1 {
			break
		}
	}
}
