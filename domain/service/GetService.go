package service

import (
	"poc-localstack-dynamo-golang/domain/entities"
	"poc-localstack-dynamo-golang/domain/repository"
)

type GetService interface {
	Get(parameter entities.Parameter) (entities.Parameter,error)
}

type GetServiceImpl struct {

}

func NewGetService() GetService {
	return &GetServiceImpl{}
}

func (s GetServiceImpl) Get(parameter entities.Parameter) (entities.Parameter, error) {

	managementRepository := repository.NewParameterRepository()

	parameterFound, err := managementRepository.Get(parameter)

	if err != nil{
		return entities.Parameter{}, err
	}

	return parameterFound, err
}