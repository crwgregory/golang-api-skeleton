package errors

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"net/http"
)

type DynamoPutError struct {
	ApiError
}

func (d DynamoPutError) Error() string {
	if d.Message == "" {
		return "there was an error creating the resource in dynamo"
	}
	return d.Message
}
func (d DynamoPutError) GetStatusCode() int {
	if d.StatusCode == 0 {
		return http.StatusInternalServerError
	}
	return d.StatusCode
}

type DynamoQueryError struct {
	ApiError
	Query *dynamodb.QueryInput
}

func (d DynamoQueryError) Error() string {
	if d.Message == "" {
		return "there was an error querying dynamodb"
	}
	return d.Message
}
func (d DynamoQueryError) GetStatusCode() int {
	if d.StatusCode == 0 {
		return http.StatusInternalServerError
	}
	return d.StatusCode
}
