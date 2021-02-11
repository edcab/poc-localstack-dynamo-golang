package dynamodb

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"poc-localstack-dynamo-golang/infrastructure/dynamodb/model"
)

type DynamoDB interface {
	Save(item model.Item) error
	Get(item model.Item) (model.Item, error)
}

type DynamoService struct {
	dynamoDBAPI dynamodbiface.DynamoDBAPI
	tableName   string
}

func NewDynamoClient(dynamoConfig dynamodbiface.DynamoDBAPI, tableName string) DynamoDB {
	return &DynamoService{
		dynamoDBAPI: dynamoConfig,
		tableName:   tableName,
	}
}

func (d DynamoService) Save(item model.Item) error {

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
	_, err := d.dynamoDBAPI.PutItem(input)

	if err != nil {
		return err
	}

	return nil
}

func (d DynamoService) Get(item model.Item) (model.Item, error) {

	result, err := d.dynamoDBAPI.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(d.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Key": {
				S: aws.String(item.Key),
			},
		},
	})

	if err != nil {
		return model.Item{}, errors.New("error getting device from DB, error: " + err.Error())
	}

	if len(result.Item) == 0 {
		return model.Item{}, errors.New("item not found")
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	return item, nil
}
