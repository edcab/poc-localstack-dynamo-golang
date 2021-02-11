package usecase

import (
	"poc-localstack-dynamo-golang/domain/dto"
	"poc-localstack-dynamo-golang/domain/entities"
	"poc-localstack-dynamo-golang/domain/service"
)

type Get interface {
	Get(dto dto.ParameterDTO) (dto.ParameterDTO, error)
}

type GetImpl struct {
}

func NewUseCaseGet() Get {
	return &GetImpl{
	}
}

func (s GetImpl) Get(parameterDto dto.ParameterDTO) (dto.ParameterDTO,error) {
	getService := service.NewGetService()

	parameter := entities.Parameter{
		Key:   parameterDto.Key,
		Value: parameterDto.Value,
	}

	entParameterFound, err := getService.Get(parameter)

	if err != nil {
		return dto.ParameterDTO{}, err
	}

	parameterDto.Value = entParameterFound.Value

	return parameterDto, err
}