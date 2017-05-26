package records

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"reflect"
)

type DynamoRecord struct {
	Record
}

func (d DynamoRecord) GetConnection() connection.ConnectionInterface {
	if d.connection == nil {
		d.connection = new(connection.DynamoDB)
	}
	return d.connection
}

func (d DynamoRecord) GetDb() *dynamodb.DynamoDB {
	return d.GetConnection().GetDB().(*dynamodb.DynamoDB)
}

func GetDynamoAttribute(value *dynamodb.AttributeValue, kind reflect.Kind) (data []byte, err error) {

	switch kind {
	case reflect.Int:
		return []byte(*value.N), nil
	case reflect.String:
		return []byte(*value.S), nil
	}
	return nil, nil
}
