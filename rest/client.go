package rest

import (
	"github.com/cksidharthan/go-bybit/rest/linear"
	"github.com/cksidharthan/go-bybit/rest/spot"
)

type Interface interface {
	Spot() spot.Interface
	Linear() linear.Interface
}

type Client struct {
	spot   spot.Interface
	linear linear.Interface
}

func (c *Client) Spot() spot.Interface {
	return c.spot
}

func (c *Client) Linear() linear.Interface {
	return c.linear
}

func NewRestClient(url, apiKey, apiSecret string) Interface {
	return &Client{
		spot:   spot.NewSpotClient(url, apiKey, apiSecret),
		linear: linear.NewLinearClient(url, apiKey, apiSecret),
	}
}
