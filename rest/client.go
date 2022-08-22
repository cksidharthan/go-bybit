package rest

import (
	"github.com/cksidharthan/go-bybit/rest/inverse_perpetual"
	"github.com/cksidharthan/go-bybit/rest/linear"
	"github.com/cksidharthan/go-bybit/rest/spot"
)

type Client struct {
	spot             spot.Client
	linear           linear.Client
	inversePerpetual inverse_perpetual.Client
}

func (c *Client) Spot() spot.Client {
	return c.spot
}

func (c *Client) Linear() linear.Client {
	return c.linear
}

func (c *Client) InversePerpetual() inverse_perpetual.Client {
	return c.inversePerpetual
}

func NewRestClient(url, apiKey, apiSecret string) *Client {
	return &Client{
		spot:             *spot.NewSpotClient(url, apiKey, apiSecret),
		linear:           *linear.NewLinearClient(url, apiKey, apiSecret),
		inversePerpetual: *inverse_perpetual.NewInversePerpetualClient(url, apiKey, apiSecret),
	}
}
