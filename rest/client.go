package rest

import (
	"github.com/cksidharthan/go-bybit/rest/linear"
	"github.com/cksidharthan/go-bybit/rest/spot"
)

type Client struct {
	spot   spot.Client
	linear linear.Client
}

func (c *Client) Spot() spot.Client {
	return c.spot
}

func (c *Client) Linear() linear.Client {
	return c.linear
}

func NewRestClient(url, apiKey, apiSecret string) *Client {
	return &Client{
		spot:   *spot.NewSpotClient(url, apiKey, apiSecret),
		linear: *linear.NewLinearClient(url, apiKey, apiSecret),
	}
}
