package rest

import "github.com/cksidharthan/go-bybit/rest/spot"

type Interface interface {
	Spot() spot.SpotInterface
}

type Client struct {
	spot spot.SpotInterface
}

func (c *Client) Spot() spot.SpotInterface {
	return c.spot
}

func New(url, apiKey, apiSecret string) Interface {
	return &Client{
		spot: spot.New(url, apiKey, apiSecret),
	}
}
