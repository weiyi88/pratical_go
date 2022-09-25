package main

import (
	"gin_demo/go-kit/endpoint"
	"gin_demo/go-kit/service"
	"gin_demo/go-kit/transport"
	"net/http"

	httpTransport "github.com/go-kit/kit/transport/http"
)

func main() {
	s := service.NewService()

	add := httpTransport.NewServer(
		endpoint.MakeAddEndpoint(s),
		transport.DecodeRequest,
		transport.EncodeResponse,
	)

	reduce := httpTransport.NewServer(
		endpoint.MakeReduceEndpoint(s),
		transport.DecodeRequest,
		transport.EncodeResponse,
	)

	multi := httpTransport.NewServer(
		endpoint.MakeMultiEndpoint(s),
		transport.DecodeRequest,
		transport.EncodeResponse,
	)
	http.Handle("/add", add)
	http.Handle("/reduce", reduce)
	http.Handle("/multi", multi)
	http.ListenAndServe(":9009", nil)
}
