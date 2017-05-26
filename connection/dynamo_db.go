package connection

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDB struct {
	Session *session.Session
}

func (d *DynamoDB) init() {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1"), CredentialsChainVerboseErrors: aws.Bool(true)})

	if err != nil {
		panic(err)
	}
	d.Session = sess
}

func (d *DynamoDB) GetDB() interface{} {

	if d.Session == nil {
		d.init()
	}
	return dynamodb.New(d.Session)
}
