package spot

import "github.com/cksidharthan/go-bybit/rest/spot/market"

type Interface interface {
	Market() market.Interface
}

type Client struct {
	market market.Interface
}

func (c *Client) Market() market.Interface {
	return c.market
}

func New(url, apiKey, apiSecret string) Interface {
	return &Client{
		market: market.New(url, apiKey, apiSecret),
	}
}
