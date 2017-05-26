package components

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/crwgregory/golang-api-skeleton/config"
	"github.com/crwgregory/golang-api-skeleton/connection"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Log struct {
	HandlerName string
	RouteName   string
	Request     http.Request
	When        time.Time
	Response    ApiResponse
}

// LogRequest logs the request to dynamodb on aws
func LogRequest(logChan chan Log) {

	dynamoCon := new(connection.DynamoDB)
	dynamo := dynamoCon.GetDB().(*dynamodb.DynamoDB)

	for {
		time.Sleep(1 * time.Second) // wait to get more logs in the channel
		logStruct := <-logChan      // block until there is at least one item

		var logPool []Log
		logPool = append(logPool, logStruct)
		for {
			if len(logPool) > config.LogPoolSize || len(logChan) == 0 {
				break
			}
			l := <-logChan
			logPool = append(logPool, l)
		}

		reqItems := make(map[string]*dynamodb.PutRequest)

		for _, l := range logPool {
			req := l.Request
			nowStr := strconv.Itoa(int(l.When.Unix()))

			log.Println(nowStr, req.RemoteAddr, l.HandlerName, l.RouteName, req.Method, req.URL.Path)

			putRequest := &dynamodb.PutRequest{

				Item: map[string]*dynamodb.AttributeValue{

					"request_path": { // primary key
						S: aws.String(req.URL.Path),
					},
					"when": { // primary key
						N: aws.String(nowStr),
					},
					"handler": {
						S: aws.String(l.HandlerName),
					},
					"route_name": {
						S: aws.String(l.RouteName),
					},
					"remote": {
						S: aws.String(req.RemoteAddr),
					},
					"method": {
						S: aws.String(req.Method),
					},
					"message": {
						S: aws.String(GetResponseMessage(l.Response)),
					},
					"res_status_code": {
						N: aws.String(strconv.Itoa(l.Response.StatusCode)),
					},
				},
			}

			key := req.URL.Path + nowStr // make a key of table primary keys so that we don't send duplicates to dynamo
			reqItems[key] = putRequest
		}

		var writeReq []*dynamodb.WriteRequest

		for _, req := range reqItems {
			writeReq = append(writeReq, &dynamodb.WriteRequest{
				PutRequest: req,
			})
		}

		params := &dynamodb.BatchWriteItemInput{
			RequestItems: map[string][]*dynamodb.WriteRequest{
				"ApiLog": writeReq,
			},
		}

		_, err := dynamo.BatchWriteItem(params)

		if err != nil {
			panic(err)
		}
	}
}
