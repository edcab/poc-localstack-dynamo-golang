package save

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"net/http"
	"poc-localstack-dynamo-golang/domain/dto"
	"poc-localstack-dynamo-golang/infrastructure/http/rest/handler/save/model"
	"poc-localstack-dynamo-golang/ports"
)

func MakeSaveEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var body model.Request
		req := request.(*http.Request)

		err := json.NewDecoder(req.Body).Decode(&body)
		if err != nil {
			return nil, err
		}

		savePort := ports.NewSavePort()

		parameterDTO := dto.ParameterDTO{
			Key:   body.Key,
			Value: body.Value,
		}

		err = savePort.Save(parameterDTO)

		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}
