package http

import (
	"context"

	"github.com/sony/gobreaker"
)

type httpClient interface {
	Post(ctx context.Context, uri string, body string, header map[string]string) (string, error)
	Get(ctx context.Context, uri string, header map[string]string) (string, error)
}

type httpCircuitBreaker struct {
	cb         *gobreaker.CircuitBreaker
	httpClient httpClient
}

func NewHttpCircuitBreaker(
	cb *gobreaker.CircuitBreaker,
	httpClient httpClient,
) *httpCircuitBreaker {
	return &httpCircuitBreaker{
		cb:         cb,
		httpClient: httpClient,
	}
}

func (hcb *httpCircuitBreaker) Post(ctx context.Context, uri string, body string, header map[string]string) (string, error) {
	resp, err := hcb.cb.Execute(func() (interface{}, error) {
		return hcb.httpClient.Post(ctx, uri, body, header)
	})
	if resp == nil {
		return "", err
	}
	return resp.(string), err
}

func (hcb *httpCircuitBreaker) Get(ctx context.Context, uri string, header map[string]string) (string, error) {
	resp, err := hcb.cb.Execute(func() (interface{}, error) {
		return hcb.httpClient.Get(ctx, uri, header)
	})
	return resp.(string), err
}
