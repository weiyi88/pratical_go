package endpoint

import (
	"context"
	"fmt"
	"gin_demo/go-kit/service"
	"github.com/go-kit/kit/endpoint"
)

type Request struct {
	A int `json:"a" form:"a"`
	B int `json:"b" form:"b"`
}

type Res struct {
	Res int   `json:"res"`
	Err error `json:"err"`
}

type Test struct {
	Msg  int `json:"message" form:"message"`
	Code int `json:"code" form:":code"`
}

func MakeAddEndpoint(s service.CalculateService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(Request)
		return Res{
			Res: s.Add(req.A, req.B),
		}, nil
	}
}

func MakeReduceEndpoint(s service.CalculateService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(Request)
		return Res{
			Res: s.Reduce(req.A, req.B),
		}, nil
	}
}
func MakeMultiEndpoint(s service.CalculateService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(Request)
		return Res{
			Res: s.Mulit(req.A, req.B),
		}, nil
	}
}

func TryTest(s service.CalculateService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(Request)
		fmt.Println(req)
		return Res{
			Res: s.Test(req.A, req.B),
		}, nil
	}
}
