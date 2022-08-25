package rest

import (
	inverseperp "github.com/cksidharthan/go-bybit/rest/inverse_perpetual"
	"github.com/cksidharthan/go-bybit/rest/linear"
	"github.com/cksidharthan/go-bybit/rest/spot"
)

type Client struct {
	spot             spot.Client
	linear           linear.Client
	inversePerpetual inverseperp.Client
}

func (c *Client) Spot() spot.Client {
	return c.spot
}

func (c *Client) Linear() linear.Client {
	return c.linear
}

func (c *Client) InversePerpetual() inverseperp.Client {
	return c.inversePerpetual
}

// NewRestClient - creates a new bybit rest client
// docs - https://bybit-exchange.github.io/docs/spot/#t-introduction
func NewRestClient(url, apiKey, apiSecret string) *Client {
	return &Client{
		spot:             *spot.NewSpotClient(url, apiKey, apiSecret),
		linear:           *linear.NewLinearClient(url, apiKey, apiSecret),
		inversePerpetual: *inverseperp.NewInversePerpetualClient(url, apiKey, apiSecret),
	}
}
