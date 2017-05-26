package main

import (
	"github.com/crwgregory/golang-api-skeleton/config"
	"github.com/crwgregory/golang-api-skeleton/handlers"
	"log"
	"net/http"
)

var routes = handlers.HandlerRoutes{
	handlers.HandlerRoute{
		Name:    "Hello",
		Path:    "/hello",
		Handler: new(handlers.HelloWorldHandler),
	},
	handlers.HandlerRoute{
		Name:    "Test",
		Path:    "/test",
		Handler: new(handlers.TestHandler),
	},
}

func main() {
	r := BuildRouter(routes)
	log.Println("starting server on port: " + config.APPLICATION_PORT)
	log.Fatal(http.ListenAndServe(":"+config.APPLICATION_PORT, r))
}
