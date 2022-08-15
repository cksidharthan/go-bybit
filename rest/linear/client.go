package linear

import (
	"github.com/cksidharthan/go-bybit/rest/domain/linear"
	"github.com/cksidharthan/go-bybit/rest/linear/account"
	"github.com/cksidharthan/go-bybit/rest/linear/market"
	"github.com/cksidharthan/go-bybit/rest/linear/wallet"
)

type Client struct {
	market  linear.MarketInterface
	account linear.AccountInterface
	wallet  linear.WalletInterface
}

func (c *Client) Market() linear.MarketInterface {
	return c.market
}

func (c *Client) Account() linear.AccountInterface {
	return c.account
}

func (c *Client) Wallet() linear.WalletInterface {
	return c.wallet
}

func NewLinearClient(url, apiKey, apiSecret string) *Client {
	return &Client{
		market:  market.NewLinearMarketClient(url, apiKey, apiSecret),
		account: account.NewLinearAccountClient(url, apiKey, apiSecret),
		wallet:  wallet.NewLinearWalletClient(url, apiKey, apiSecret),
	}
}
