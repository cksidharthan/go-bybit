package spot

import (
	accountMethod "github.com/cksidharthan/go-bybit/rest/domain/spot/account/methods"
	marketMethod "github.com/cksidharthan/go-bybit/rest/domain/spot/market/methods"
	"github.com/cksidharthan/go-bybit/rest/spot/account"
	"github.com/cksidharthan/go-bybit/rest/spot/market"
)

type Interface interface {
	Market() marketMethod.Interface
	Account() accountMethod.Interface
}

type Client struct {
	market  marketMethod.Interface
	account accountMethod.Interface
}

func (c *Client) Market() marketMethod.Interface {
	return c.market
}

func (c *Client) Account() accountMethod.Interface {
	return c.account
}

func New(url, apiKey, apiSecret string) Interface {
	return &Client{
		market:  market.New(url, apiKey, apiSecret),
		account: account.New(url, apiKey, apiSecret),
	}
}
