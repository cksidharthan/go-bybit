package gobybit

import (
	"github.com/cksidharthan/go-bybit/rest"
)

type RestClientInterface interface {
	Rest() rest.Interface
}

type RestClient struct {
	rest rest.Interface
}

func (c *RestClient) Rest() rest.Interface {
	return c.rest
}

func NewBybitClient(url, apiKey, apiSecret string) *RestClient {
	return &RestClient{
		rest: rest.NewRestClient(url, apiKey, apiSecret),
	}
}
