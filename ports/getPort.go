package ports

import (
	"poc-localstack-dynamo-golang/domain/dto"
	"poc-localstack-dynamo-golang/usecase"
)

type GetPort interface {
	Get(parameter dto.ParameterDTO ) (dto.ParameterDTO, error)
}

type GetPortImpl struct {
}

func NewGetPort() GetPort{
	return &GetPortImpl{}
}

func (s GetPortImpl) Get(parameter dto.ParameterDTO) (dto.ParameterDTO, error) {
	ucGet := usecase.NewUseCaseGet()
	parameterDtoFound, err := ucGet.Get(parameter)

	if err != nil {
		return dto.ParameterDTO{}, err
	}

	return parameterDtoFound, nil
}



