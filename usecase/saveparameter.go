package usecase

import (
	"poc-localstack-dynamo-golang/domain/dto"
	"poc-localstack-dynamo-golang/domain/entities"
	"poc-localstack-dynamo-golang/domain/service"
)

type Save interface {
	Save(dto dto.ParameterDTO) error
}

type SaveImpl struct {
}

func NewUseCaseSave() Save {
	return &SaveImpl{}
}

func (s SaveImpl) Save(dto dto.ParameterDTO) error {
	saveService := service.NewSaveService()

	parameter := entities.Parameter{
		Key:   dto.Key,
		Value: dto.Value,
	}

	err := saveService.Save(parameter)

	if err != nil {
		return err
	}

	return nil
}
