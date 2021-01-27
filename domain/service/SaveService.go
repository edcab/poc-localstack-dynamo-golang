package service

import (
	"poc-localstack-dynamo-golang/domain/entities"
	"poc-localstack-dynamo-golang/domain/repository"
)

type SaveService interface {
	Save(parameter entities.Parameter) error
}

type SaveServiceImpl struct {

}

func NewSaveService() SaveService {
	return &SaveServiceImpl{}
}

func (s SaveServiceImpl) Save(parameter entities.Parameter) error {

	managementRepository := repository.NewParameterRepository()

	err := managementRepository.Save(parameter)

	if err != nil{
		return err
	}

	return nil
}