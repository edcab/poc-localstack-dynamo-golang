package ports

import (
	"poc-localstack-dynamo-golang/domain/dto"
	"poc-localstack-dynamo-golang/usecase"
)

type SavePort interface {
	Save(parameter dto.ParameterDTO) error
}

type SavePortImpl struct {
}

func NewSavePort() SavePort {
	return &SavePortImpl{}
}

func (s SavePortImpl) Save(parameter dto.ParameterDTO) error {
	ucSave := usecase.NewUseCaseSave()
	err := ucSave.Save(parameter)

	if err != nil {
		return err
	}

	return nil
}
