package save

import (
	"context"
	"encoding/json"
	"net/http"
)

const ContentType = "Content-Type"
const ValueContentType = "application/json; charset=UTF-8"

func DecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set(ContentType, ValueContentType)
	w.WriteHeader(http.StatusNoContent)

	return json.NewEncoder(w).Encode(response)
}
