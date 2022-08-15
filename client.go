package gobybit

import (
	"github.com/cksidharthan/go-bybit/rest"
)

type RestClient struct {
	rest rest.Client
}

func (c *RestClient) Rest() rest.Client {
	return c.rest
}

func NewBybitClient(url, apiKey, apiSecret string) *RestClient {
	return &RestClient{
		rest: *rest.NewRestClient(url, apiKey, apiSecret),
	}
}
