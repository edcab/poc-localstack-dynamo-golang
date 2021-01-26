package save

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"net/http"
	"poc-localstack-dynamo-golang/infrastructure/http/rest/handler/save/model"
)

func MakeSaveEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var body model.Request
		req := request.(*http.Request)

		err := json.NewDecoder(req.Body).Decode(&body)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}
