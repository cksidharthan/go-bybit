package rest

import "github.com/cksidharthan/go-bybit/rest/spot"

type Interface interface {
	Spot() spot.Interface
}

type Client struct {
	spot spot.Interface
}

func (c *Client) Spot() spot.Interface {
	return c.spot
}

func New(url, apiKey, apiSecret string) Interface {
	return &Client{
		spot: spot.New(url, apiKey, apiSecret),
	}
}
