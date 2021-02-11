package rest

import (
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"poc-localstack-dynamo-golang/infrastructure/http/rest/handler/get"
	"poc-localstack-dynamo-golang/infrastructure/http/rest/handler/save"
)

const (
	URLSave = "/v1"
	URLGet  = "/v1"
)

type Router interface {
	CreateHandler()
}

type HTTPHandler struct {
}

func NewHandler() *HTTPHandler {
	return &HTTPHandler{}
}

type endpoints struct {
	Save func() endpoint.Endpoint
	Get  func() endpoint.Endpoint
}

func makeServerEndpoints() endpoints {
	return endpoints{
		Save: save.MakeSaveEndpoint,
		Get:  get.MakeGetEndpoint,
	}
}

func (http HTTPHandler) CreateHandler() http.Handler {

	r := mux.NewRouter()
	e := makeServerEndpoints()

	r.Methods("POST").Path(URLSave).Handler(kithttp.NewServer(
		e.Save(),
		save.DecodeRequest,
		save.EncodeResponse,
	))
	r.Methods("GET").Path(URLGet).Handler(kithttp.NewServer(
		e.Get(),
		get.DecodeRequest,
		get.EncodeResponse,
	))

	return r
}
