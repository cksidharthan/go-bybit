package spot

import (
	"github.com/cksidharthan/go-bybit/rest/domain/spot"
	"github.com/cksidharthan/go-bybit/rest/spot/account"
	"github.com/cksidharthan/go-bybit/rest/spot/market"
)

type Interface interface {
	Market() spot.MarketInterface
	Account() spot.AccountInterface
}

type Client struct {
	market  spot.MarketInterface
	account spot.AccountInterface
}

func (c *Client) Market() spot.MarketInterface {
	return c.market
}

func (c *Client) Account() spot.AccountInterface {
	return c.account
}

func New(url, apiKey, apiSecret string) Interface {
	return &Client{
		market:  market.New(url, apiKey, apiSecret),
		account: account.New(url, apiKey, apiSecret),
	}
}
