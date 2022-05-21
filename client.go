package go_bybit

import (
	"github.com/cksidharthan/go-bybit/rest"
)

type BybitClientInterface interface {
	Rest() rest.Interface
}

type BybitClient struct {
	rest rest.Interface
}

func (c *BybitClient) Bybit() rest.Interface {
	return c.rest
}

func New(url, apiKey, apiSecret string) *BybitClient {
	return &BybitClient{
		rest: rest.New(url, apiKey, apiSecret),
	}
}
