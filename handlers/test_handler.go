package handlers

import (
	"fmt"
	"github.com/crwgregory/golang-api-skeleton/components"
	"math/rand"
	"time"
)

type TestHandler struct{}

func (e *TestHandler) Handle(r Request, logChan chan components.Log) components.ApiResponse {
	// heavy load
	sleep := time.Duration(rand.Intn(2)) * time.Second
	time.Sleep(sleep)
	return components.ApiResponse{
		StatusCode: 200,
		Message:    fmt.Sprintf("slept for %s", sleep.String()),
	}
}
