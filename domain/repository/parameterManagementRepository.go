package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	awsDynamo "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"poc-localstack-dynamo-golang/domain/entities"
	"poc-localstack-dynamo-golang/infrastructure/dynamodb"
	"poc-localstack-dynamo-golang/infrastructure/dynamodb/model"
)

type ParameterManagementRepository interface {
	Save(parameter entities.Parameter) error
}

type ParameterManagementRepositoryImpl struct {
}

func NewParameterRepository() ParameterManagementRepository {
	return &ParameterManagementRepositoryImpl{
	}
}

func (p ParameterManagementRepositoryImpl) Save(parameter entities.Parameter) error {

	dynamoClient := dynamodb.NewDynamoClient(NewSession("us-east-1"), "ParameterAPI")

	itemToSave := model.Item{
		Key:   parameter.Key,
		Value: parameter.Value,
	}

	err := dynamoClient.Save(itemToSave)

	if err != nil{
		return err
	}

	return nil

}


func NewSession(region string) dynamodbiface.DynamoDBAPI {
	endpoint := "http://localhost:4566"

	// Create Dynamodb AWS session
	config := &aws.Config{
		Endpoint: &endpoint,
		Region:   &region,
	}

	var svc dynamodbiface.DynamoDBAPI
	sess := session.Must(session.NewSession(config))
	svc = awsDynamo.New(sess)

	return svc
}
