package get

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"net/http"
	"poc-localstack-dynamo-golang/domain/dto"
	"poc-localstack-dynamo-golang/infrastructure/http/rest/handler/get/model"
	"poc-localstack-dynamo-golang/ports"
)

func MakeGetEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var body model.Request
		req := request.(*http.Request)

		err := json.NewDecoder(req.Body).Decode(&body)
		if err != nil {
			return nil, err
		}

		port := ports.NewGetPort()

		parameterDTO := dto.ParameterDTO{
			Key: body.Key,
		}

		parameterFoundDTO, err := port.Get(parameterDTO)

		if err != nil {
			return model.Response{}, err
		}

		response := model.Response{Value: parameterFoundDTO.Value}

		return response, nil
	}
}
