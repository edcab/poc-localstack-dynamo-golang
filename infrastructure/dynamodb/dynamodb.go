package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"poc-localstack-dynamo-golang/infrastructure/dynamodb/model"
)

type DynamoDB interface{
	Save(item model.Item) error
	Get() (model.Item, error)
}

type DynamoService struct {
	dynamoConfig dynamodbiface.DynamoDBAPI
	tableName    string
}

func NewDynamoClient(dynamoConfig dynamodbiface.DynamoDBAPI, tableName string) DynamoDB {
	return &DynamoService{
		dynamoConfig: dynamoConfig,
		tableName:    tableName,
	}
}

func (d DynamoService) Save(item model.Item) error{

	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"Key": {
				S: aws.String(item.Key),
			}, "Value": {
				S: aws.String(item.Value),
			},
		},
		TableName: aws.String(d.tableName),
	}


	_, err := d.dynamoConfig.PutItem(input)

	if err != nil {
		return err
	}

	return nil
}

func (d DynamoService) Get() (model.Item,error) {
	return model.Item{}, nil
}