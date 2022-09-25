package transport

import (
	"context"
	"encoding/json"
	"errors"
	"gin_demo/go-kit/endpoint"
	"net/http"
)

func DecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, errors.New("params error")
	}
	return req, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, res interface{}) error {
	return json.NewEncoder(w).Encode(res)
}
